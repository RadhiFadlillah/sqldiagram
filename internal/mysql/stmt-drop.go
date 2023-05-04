package mysql

import (
	"github.com/RadhiFadlillah/sqldiagram/internal/common"
	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
)

func parseDropTable(stmt *ast.DropTableStmt, current []Table) []Table {
	deletedTables := common.NewSet[string]()
	for _, table := range stmt.Tables {
		deletedTables.Put(table.Name.String())
	}

	var leftoverTables []Table
	for _, table := range current {
		if !deletedTables.Has(table.Name) {
			leftoverTables = append(leftoverTables, table)
		}
	}

	return leftoverTables
}
