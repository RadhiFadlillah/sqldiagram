package mysql

import (
	"testing"

	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/stretchr/testify/assert"
)

func Test_parseDropTable(t *testing.T) {
	var node ast.StmtNode
	var stmtDrop *ast.DropTableStmt
	tables := []Table{
		{Name: "person", Columns: []Column{
			{Name: "id", Type: "INT"},
			{Name: "name", Type: "VARCHAR"},
			{Name: "age", Type: "INT"},
			{Name: "zodiac", Type: "VARCHAR"},
		}},
	}

	// Remove unexisting table
	node = parseSingleQuery(`DROP TABLE company`)
	stmtDrop = node.(*ast.DropTableStmt)
	tables = parseDropTable(stmtDrop, tables)
	assert.Len(t, tables, 1)

	// Remove existing table
	node = parseSingleQuery(`DROP TABLE person`)
	stmtDrop = node.(*ast.DropTableStmt)
	tables = parseDropTable(stmtDrop, tables)
	assert.Len(t, tables, 0)
}
