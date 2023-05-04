package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"github.com/RadhiFadlillah/sqldiagram/internal/mysql"
	"github.com/pingcap/tidb/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/urfave/cli/v2"
)

func cmdMySql() *cli.Command {
	cmd := &cli.Command{
		Name:   "mysql",
		Usage:  "generate ERD from MySQL dialect",
		Action: cmdMySqlAction,
	}

	return cmd
}

func cmdMySqlAction(ctx *cli.Context) error {
	// Parse flags
	dontUseGroup := ctx.Bool(appNoGroup)
	renderRawD2 := ctx.Bool(appRawD2)
	outputPath := ctx.String(appOutput)

	diagramDirection := ctx.String(appDirection)
	diagramDirection = strings.ToLower(diagramDirection)
	switch diagramDirection {
	case "up", "down", "left":
	default:
		diagramDirection = "right"
	}

	// Get input path from args
	inputPaths := ctx.Args().Slice()
	if len(inputPaths) == 0 {
		inputPaths = []string{"."}
	}

	sqlFiles, err := getSqlFiles(inputPaths...)
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

	// If necessary, attach group name to table relations
	if len(groups) > 1 {
		// Map table name to its group
		mapTableGroup := map[string]string{}
		for _, g := range groups {
			for _, t := range g.Tables {
				mapTableGroup[t.Name] = g.Name
			}
		}

		// Put group name to related tables
		for _, g := range groups {
			for _, t := range g.Tables {
				for i, c := range t.Columns {
					for i, ref := range c.ReferTo {
						if refGroup, exist := mapTableGroup[ref]; exist {
							c.ReferTo[i] = refGroup + "." + ref
						}
					}
					t.Columns[i] = c
				}
			}
		}
	}

	// Generate d2 diagram and graph
	diagram, graph, err := generateD2(groups, diagramDirection)
	if err != nil {
		return err
	}

	// Render d2 to preferred format
	var renderResult []byte
	if renderRawD2 {
		renderResult = renderScript(graph)
	} else {
		renderResult, err = renderSvg(diagram)
		if err != nil {
			return err
		}
	}

	// Write the render result
	return writeOutput(renderResult, outputPath, renderRawD2)
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
	tables := mysql.Parse(queries, nil)

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
