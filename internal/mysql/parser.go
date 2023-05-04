package mysql

import (
	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
)

func Parse(stmts []ast.StmtNode, current []Table) []Table {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.CreateTableStmt:
			current = parseCreateTable(s, current)
		case *ast.AlterTableStmt:
			current = parseAlterTable(s, current)
		}
	}

	return current
}
