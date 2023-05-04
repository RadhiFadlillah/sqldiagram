package mysql

import (
	"testing"

	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/stretchr/testify/assert"
)

var sqlParser = parser.New()

func Test_parseCreateTable(t *testing.T) {
	var node ast.StmtNode
	var stmtTable *ast.CreateTableStmt
	var table Table
	var tables []Table
	var columns []Column

	// Basic create query
	node = parseSingleQuery(`
	CREATE TABLE person (
		id   INT          NOT NULL,
		name VARCHAR(255) NOT NULL
	)`)

	stmtTable = node.(*ast.CreateTableStmt)
	tables = parseCreateTable(stmtTable, nil)
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 2)

	assert.Equal(t, "id", columns[0].Name)
	assert.Equal(t, "INT", columns[0].Type)

	assert.Equal(t, "name", columns[1].Name)
	assert.Equal(t, "VARCHAR", columns[1].Type)

	// Create query with inline keys
	node = parseSingleQuery(`
	CREATE TABLE person (
		id         INT           NOT NULL PRIMARY KEY,
		company_id INT           NOT NULL REFERENCES company (id),
		identifier VARBINARY(20) NOT NULL UNIQUE KEY,
		name       VARCHAR(255)  NOT NULL
	)`)

	stmtTable = node.(*ast.CreateTableStmt)
	tables = parseCreateTable(stmtTable, nil)
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 4)

	assert.Equal(t, "id", columns[0].Name)
	assert.Equal(t, "INT", columns[0].Type)
	assert.True(t, columns[0].IsPK)

	assert.Equal(t, "company_id", columns[1].Name)
	assert.Equal(t, "INT", columns[1].Type)
	assert.True(t, columns[1].IsFK)
	assert.Len(t, columns[1].ReferTo, 1)
	assert.Equal(t, "company", columns[1].ReferTo[0])

	assert.Equal(t, "identifier", columns[2].Name)
	assert.Equal(t, "VARBINARY", columns[2].Type)
	assert.True(t, columns[2].IsUnique)

	assert.Equal(t, "name", columns[3].Name)
	assert.Equal(t, "VARCHAR", columns[3].Type)

	// Create query with separate constraint
	node = parseSingleQuery(`
	CREATE TABLE person (
		id         INT           NOT NULL,
		company_id INT           NOT NULL,
		identifier VARBINARY(20) NOT NULL,
		name       VARCHAR(255)  NOT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY person_identifier_UQ (identifier),
		FOREIGN KEY person_company_id_FK (company_id) REFERENCES company (id)
	)`)

	stmtTable = node.(*ast.CreateTableStmt)
	tables = parseCreateTable(stmtTable, nil)
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 4)

	assert.Equal(t, "id", columns[0].Name)
	assert.Equal(t, "INT", columns[0].Type)
	assert.True(t, columns[0].IsPK)

	assert.Equal(t, "company_id", columns[1].Name)
	assert.Equal(t, "INT", columns[1].Type)
	assert.True(t, columns[1].IsFK)
	assert.Len(t, columns[1].ReferTo, 1)
	assert.Equal(t, "company", columns[1].ReferTo[0])

	assert.Equal(t, "identifier", columns[2].Name)
	assert.Equal(t, "VARBINARY", columns[2].Type)
	assert.True(t, columns[2].IsUnique)

	assert.Equal(t, "name", columns[3].Name)
	assert.Equal(t, "VARCHAR", columns[3].Type)
}
