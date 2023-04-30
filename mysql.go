package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/urfave/cli/v2"
)

var rxColumnType = regexp.MustCompile(`(?i)^[a-z]+`)

func cmdMySql() *cli.Command {
	cmd := &cli.Command{
		Name:   "mysql",
		Usage:  "generate ERD from MySQL dialect",
		Action: cmdMySqlAction,
	}

	return cmd
}

func cmdMySqlAction(ctx *cli.Context) error {
	// Get input dir from args
	inputDir := ctx.Args().Get(0)
	if inputDir == "" {
		inputDir = "."
	}

	// Parse flags
	dontUseGroup := ctx.Bool(appFlNoGroup)

	// Catch SQL files inside input dir
	sqlFiles, err := getSqlFiles(inputDir)
	if err != nil {
		return err
	}

	// Parse and extract tables from each file
	p := parser.New()
	var groups []Group
	for _, sqlFile := range sqlFiles {
		// Extract group from file
		group, err := parseMySqlFile(p, sqlFile)
		if err != nil {
			return err
		}

		// Save group
		if group != nil && len(group.Tables) > 0 {
			groups = append(groups, *group)
		}
	}

	// Make sure there is a table found
	if len(groups) == 0 {
		return fmt.Errorf("no table found")
	}

	// If necessary, merge all groups into one
	if len(groups) > 1 && dontUseGroup {
		root := Group{Name: "root"}
		for _, g := range groups {
			root.Tables = append(root.Tables, g.Tables...)
		}
		groups = []Group{root}
	}

	// Generate d2 codes
	d2codes, err := generateD2Codes(groups)
	if err != nil {
		return err
	}

	// Render D2 to SVG
	svgCodes, err := renderD2Svg(d2codes)
	if err != nil {
		return err
	}

	fmt.Println(string(svgCodes))
	return nil
}

func parseMySqlFile(p *parser.Parser, path string) (*Group, error) {
	// Open file
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %v", path, err)
	}

	// Extract queries
	queries, _, err := p.Parse(string(f), "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to parse %q: %v", path, err)
	}

	// Extract DDL queries only
	var tables []Table
	for _, query := range queries {
		if ddlQuery, isDDL := query.(*ast.CreateTableStmt); isDDL {
			// Extract constraints
			uniqueKeys := NewSet[string]()
			primaryKeys := NewSet[string]()
			foreignKeys := NewSet[string]()
			relatedTables := NewSet[string]()

			for _, c := range ddlQuery.Constraints {
				switch c.Tp {
				case ast.ConstraintPrimaryKey:
					for _, col := range c.Keys {
						primaryKeys.Put(col.Column.Name.String())
					}

				case ast.ConstraintUniq, ast.ConstraintUniqKey:
					for _, col := range c.Keys {
						uniqueKeys.Put(col.Column.Name.String())
					}

				case ast.ConstraintForeignKey:
					if r := c.Refer; r != nil {
						for _, col := range c.Keys {
							colName := col.Column.Name.String()
							foreignKeys.Put(colName)
						}

						dstTable := r.Table.Name.String()
						relatedTables.Put(dstTable)
					}
				}
			}

			// Extract columns
			var primaryColumns, foreignColumns, columns []Column
			for _, c := range ddlQuery.Cols {
				colName := c.Name.String()
				colType := rxColumnType.FindString(c.Tp.String())

				column := Column{
					Name:   colName,
					Tp:     strings.ToUpper(colType),
					Unique: uniqueKeys.Has(colName),
				}

				if primaryKeys.Has(colName) {
					primaryColumns = append(primaryColumns, column)
				} else if foreignKeys.Has(colName) {
					foreignColumns = append(foreignColumns, column)
				} else {
					columns = append(columns, column)
				}
			}

			// Save tables
			tables = append(tables, Table{
				Name:          ddlQuery.Table.Name.String(),
				PrimaryKeys:   primaryColumns,
				ForeignKeys:   foreignColumns,
				Columns:       columns,
				RelatedTables: relatedTables.Keys(),
			})
		}
	}

	// Prepare group data
	groupName := filepath.Base(path)
	groupName = strings.TrimSuffix(groupName, filepath.Ext(groupName))
	groupLabel := strings.ReplaceAll(groupName, "_", "-")
	groupLabel = strings.ToUpper(groupLabel)

	return &Group{
		Name:   groupName,
		Label:  groupLabel,
		Tables: tables,
	}, nil
}
