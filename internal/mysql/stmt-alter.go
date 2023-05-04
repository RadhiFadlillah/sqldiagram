package mysql

import (
	"strings"

	"github.com/RadhiFadlillah/sqldiagram/internal/common"
	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
)

func parseAlterTable(stmt *ast.AlterTableStmt, current []Table) []Table {
	// Get the altered table
	tableName := stmt.Table.Name.String()

	// Find altered table in current list
	alteredTable, idx := common.FindSlice(current,
		func(t Table) bool { return t.Name == tableName })
	if idx == -1 {
		return current
	}

	// Process each alter specs
	for _, spec := range stmt.Specs {
		switch spec.Tp {
		case ast.AlterTableDropColumn:
			oldName := spec.OldColumnName.OrigColName()
			_, oldIdx := common.FindSlice(alteredTable.Columns, func(c Column) bool { return c.Name == oldName })
			if oldIdx >= 0 {
				alteredTable.Columns = common.DeleteSliceItem(alteredTable.Columns, oldIdx)
			}

		case ast.AlterTableAddColumns:
			for _, c := range spec.NewColumns {
				colName := c.Name.OrigColName()
				colType := rxColumnType.FindString(c.Tp.String())

				var referTo []string
				var isPK, isFK, isUnique bool
				for _, opt := range c.Options {
					isPK = isPK || opt.Tp == ast.ColumnOptionPrimaryKey
					isUnique = isUnique || opt.Tp == ast.ColumnOptionUniqKey

					if opt.Refer != nil {
						isFK = true
						referTo = append(referTo, opt.Refer.Table.Name.String())
					}
				}

				alteredTable.Columns = append(alteredTable.Columns, Column{
					Name:     colName,
					Type:     strings.ToUpper(colType),
					IsPK:     isPK,
					IsFK:     isFK,
					IsUnique: isUnique,
					ReferTo:  common.UniqueSlice(referTo),
				})
			}

		case ast.AlterTableRenameColumn:
			oldName := spec.OldColumnName.OrigColName()
			_, oldIdx := common.FindSlice(alteredTable.Columns, func(c Column) bool { return c.Name == oldName })
			if oldIdx >= 0 {
				newName := spec.NewColumnName.OrigColName()
				alteredTable.Columns[oldIdx].Name = newName
			}

		case ast.AlterTableChangeColumn:
			oldName := spec.OldColumnName.OrigColName()
			_, oldIdx := common.FindSlice(alteredTable.Columns, func(c Column) bool { return c.Name == oldName })
			if oldIdx >= 0 && len(spec.NewColumns) == 1 {
				c := spec.NewColumns[0]
				colName := c.Name.OrigColName()
				colType := rxColumnType.FindString(c.Tp.String())

				var referTo []string
				var isPK, isFK, isUnique bool
				for _, opt := range c.Options {
					isPK = isPK || opt.Tp == ast.ColumnOptionPrimaryKey
					isUnique = isUnique || opt.Tp == ast.ColumnOptionUniqKey

					if opt.Refer != nil {
						isFK = true
						referTo = append(referTo, opt.Refer.Table.Name.String())
					}
				}

				alteredTable.Columns[oldIdx] = Column{
					Name:     colName,
					Type:     strings.ToUpper(colType),
					IsPK:     isPK,
					IsFK:     isFK,
					IsUnique: isUnique,
					ReferTo:  common.UniqueSlice(referTo),
				}
			}

		case ast.AlterTableAddConstraint:
			for _, key := range spec.Constraint.Keys {
				colName := key.Column.OrigColName()
				col, colIdx := common.FindSlice(alteredTable.Columns, func(c Column) bool { return c.Name == colName })
				if colIdx >= 0 {
					switch spec.Constraint.Tp {
					case ast.ConstraintPrimaryKey:
						col.IsPK = true
					case ast.ConstraintUniq, ast.ConstraintUniqKey:
						col.IsUnique = true
					case ast.ConstraintForeignKey:
						if r := spec.Constraint.Refer; r != nil {
							dstTable := r.Table.Name.String()
							col.IsFK = true
							col.ReferTo = append(col.ReferTo, dstTable)
						}
					}

					alteredTable.Columns[colIdx] = col
				}
			}
		}
	}

	current[idx] = alteredTable
	return current
}
