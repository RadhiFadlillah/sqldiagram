package mysql

import (
	"fmt"
	"reflect"

	"github.com/pingcap/tidb/parser/ast"
)

func parseSingleQuery(s string) ast.StmtNode {
	node, err := sqlParser.ParseOneStmt(s, "utf8mb4", "")
	if err != nil {
		panic(err)
	}

	switch stmt := node.(type) {
	case *ast.CreateTableStmt:
		return stmt
	case *ast.AlterTableStmt:
		return stmt
	default:
		panic(fmt.Errorf("unknown stmt type: %q", reflect.TypeOf(node)))
	}
}
