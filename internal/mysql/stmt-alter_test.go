package mysql

import (
	"testing"

	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/stretchr/testify/assert"
)

func Test_parseAlterTable(t *testing.T) {
	var node ast.StmtNode
	var stmtAlter *ast.AlterTableStmt
	var table Table
	var tables []Table
	var columns []Column

	// Drop columns
	table = Table{Name: "person", Columns: []Column{
		{Name: "id", Type: "INT"},
		{Name: "name", Type: "VARCHAR"},
		{Name: "age", Type: "INT"},
		{Name: "zodiac", Type: "VARCHAR"},
	}}

	node = parseSingleQuery(`
		ALTER TABLE person
		DROP COLUMN age,
		DROP COLUMN zodiac;`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, []Table{table})
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 2)

	assert.Equal(t, "id", columns[0].Name)
	assert.Equal(t, "INT", columns[0].Type)

	assert.Equal(t, "name", columns[1].Name)
	assert.Equal(t, "VARCHAR", columns[1].Type)

	// Add columns, simple
	table = Table{Name: "person", Columns: []Column{
		{Name: "id", Type: "INT"},
		{Name: "name", Type: "VARCHAR"},
	}}

	node = parseSingleQuery(`
		ALTER TABLE person
		ADD COLUMN age INT,
		ADD COLUMN zodiac VARCHAR(255);`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, []Table{table})
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 4)

	assert.Equal(t, "age", columns[2].Name)
	assert.Equal(t, "INT", columns[2].Type)

	assert.Equal(t, "zodiac", columns[3].Name)
	assert.Equal(t, "VARCHAR", columns[3].Type)

	// Add column as primary key
	table = Table{Name: "person", Columns: []Column{
		{Name: "name", Type: "VARCHAR"},
	}}

	node = parseSingleQuery(`
		ALTER TABLE person
		ADD COLUMN id INT PRIMARY KEY`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, []Table{table})
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 2)

	assert.Equal(t, "name", columns[0].Name)
	assert.Equal(t, "VARCHAR", columns[0].Type)

	assert.Equal(t, "id", columns[1].Name)
	assert.Equal(t, "INT", columns[1].Type)
	assert.True(t, columns[1].IsPK)

	// Change column
	table = Table{Name: "person", Columns: []Column{
		{Name: "id", Type: "INT"},
		{Name: "fake", Type: "INT"},
		{Name: "age", Type: "INT"},
	}}

	node = parseSingleQuery(`
		ALTER TABLE person
		CHANGE COLUMN fake name VARCHAR(255) PRIMARY KEY`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, []Table{table})
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 3)

	assert.Equal(t, "name", columns[1].Name)
	assert.Equal(t, "VARCHAR", columns[1].Type)
	assert.True(t, columns[1].IsPK)

	// Add constraint
	table = Table{Name: "person", Columns: []Column{
		{Name: "id", Type: "INT"},
		{Name: "name", Type: "VARCHAR"},
		{Name: "identifier", Type: "VARBINARY"},
		{Name: "company_id", Type: "INT"},
	}}

	node = parseSingleQuery(`
		ALTER TABLE person
		ADD CONSTRAINT PRIMARY KEY (id),
		ADD CONSTRAINT UNIQUE KEY person_identifier_UQ (identifier)`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, []Table{table})
	assert.Len(t, tables, 1)

	table = tables[0]
	columns = table.Columns
	assert.Equal(t, "person", table.Name)
	assert.Len(t, columns, 4)

	assert.Equal(t, "id", columns[0].Name)
	assert.Equal(t, "INT", columns[0].Type)
	assert.True(t, columns[0].IsPK)

	assert.Equal(t, "name", columns[1].Name)
	assert.Equal(t, "VARCHAR", columns[1].Type)

	assert.Equal(t, "identifier", columns[2].Name)
	assert.Equal(t, "VARBINARY", columns[2].Type)
	assert.True(t, columns[2].IsUnique)

	assert.Equal(t, "company_id", columns[3].Name)
	assert.Equal(t, "INT", columns[3].Type)

	// Continue adding constraint
	node = parseSingleQuery(`
		ALTER TABLE person
		ADD CONSTRAINT FOREIGN KEY person_company_id_FK (company_id) REFERENCES company (id)`)
	stmtAlter = node.(*ast.AlterTableStmt)
	tables = parseAlterTable(stmtAlter, tables)
	assert.Len(t, tables, 1)

	assert.True(t, columns[0].IsPK)
	assert.True(t, columns[2].IsUnique)
	assert.True(t, columns[3].IsFK)
	assert.Len(t, columns[3].ReferTo, 1)
	assert.Equal(t, "company", columns[3].ReferTo[0])

}
