package mysql

import (
	"regexp"
	"strings"

	"github.com/RadhiFadlillah/sqldiagram/internal/common"
	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
)

var rxColumnType = regexp.MustCompile(`(?i)^[a-z]+`)

func parseCreateTable(stmt *ast.CreateTableStmt, current []Table) []Table {
	// Extract constraints
	uniqueKeys := common.NewSet[string]()
	primaryKeys := common.NewSet[string]()
	foreignKeys := common.NewMap[string, []string]()

	for _, c := range stmt.Constraints {
		switch c.Tp {
		case ast.ConstraintPrimaryKey:
			for _, col := range c.Keys {
				primaryKeys.Put(col.Column.OrigColName())
			}

		case ast.ConstraintUniq, ast.ConstraintUniqKey:
			for _, col := range c.Keys {
				uniqueKeys.Put(col.Column.OrigColName())
			}

		case ast.ConstraintForeignKey:
			if r := c.Refer; r != nil {
				dstTable := r.Table.Name.String()
				for _, col := range c.Keys {
					colName := col.Column.OrigColName()
					dstTables := append(foreignKeys.Get(colName), dstTable)
					foreignKeys.Put(colName, dstTables)
				}
			}
		}
	}

	// Extract columns
	var columns []Column
	for _, c := range stmt.Cols {
		colName := c.Name.OrigColName()
		colType := rxColumnType.FindString(c.Tp.String())
		referTo := foreignKeys.Get(colName)

		var isPK, isFK, isUnique bool
		for _, opt := range c.Options {
			isPK = isPK || opt.Tp == ast.ColumnOptionPrimaryKey
			isUnique = isUnique || opt.Tp == ast.ColumnOptionUniqKey

			if opt.Refer != nil {
				isFK = true
				referTo = append(referTo, opt.Refer.Table.Name.String())
			}
		}

		columns = append(columns, Column{
			Name:     colName,
			Type:     strings.ToUpper(colType),
			IsPK:     primaryKeys.Has(colName) || isPK,
			IsFK:     foreignKeys.Has(colName) || isFK,
			IsUnique: uniqueKeys.Has(colName) || isUnique,
			ReferTo:  common.UniqueSlice(referTo),
		})
	}

	return append(current, Table{
		Name:    stmt.Table.Name.String(),
		Columns: columns,
	})
}
