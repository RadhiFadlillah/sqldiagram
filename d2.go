package main

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/RadhiFadlillah/sqldiagram/internal/common"
	. "github.com/RadhiFadlillah/sqldiagram/internal/model"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func generateD2(groups []Group, direction string) (*d2target.Diagram, *d2graph.Graph, error) {
	// Prepare template
	tpl, err := template.New("d2").
		Funcs(template.FuncMap{
			"bgColor":    getBgColor,
			"fgColor":    getFgColor,
			"column":     d2EscapeKeyword,
			"constraint": d2ColumnConstraint,
			"relations":  d2TableRelations,
		}).Parse(d2TemplateText)
	if err != nil {
		return nil, nil, fmt.Errorf("create template error: %w", err)
	}

	// Execute template to generate D2 codes
	var d2Codes []byte
	if len(groups) == 1 {
		d2Codes, err = d2CodesFromTables(tpl, groups[0].Tables, direction)
	} else {
		d2Codes, err = d2CodesFromGroups(tpl, groups, direction)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("exec template error: %w", err)
	}

	// Prepare text measurer
	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return nil, nil, fmt.Errorf("textmeasure error: %w", err)
	}

	// Generate diagram and graph
	ctx := context.Background()
	diagram, graph, err := d2lib.Compile(ctx, string(d2Codes), &d2lib.CompileOptions{
		Ruler:  ruler,
		Layout: d2elklayout.DefaultLayout})
	if err != nil {
		return nil, nil, fmt.Errorf("compile error: %w", err)
	}

	return diagram, graph, nil
}

func d2CodesFromTables(tpl *template.Template, tables []Table, direction string) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	data := d2TemplateData{Direction: direction, Tables: tables}
	if err := tpl.ExecuteTemplate(buffer, "d2Tables", data); err != nil {
		return nil, fmt.Errorf("template write error: %w", err)
	}

	return buffer.Bytes(), nil
}

func d2CodesFromGroups(tpl *template.Template, groups []Group, direction string) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	data := d2TemplateData{Direction: direction, Groups: groups}
	if err := tpl.ExecuteTemplate(buffer, "d2Groups", data); err != nil {
		return nil, fmt.Errorf("template write error: %w", err)
	}

	return buffer.Bytes(), nil
}

func d2EscapeKeyword(name string) string {
	if _, reserved := d2graph.ReservedKeywords[name]; reserved {
		return fmt.Sprintf("\"%s \"", name)
	}
	return name
}

func d2ColumnConstraint(column Column) string {
	if column.IsPK {
		return "{constraint: primary_key}"
	} else if column.IsFK {
		return "{constraint: foreign_key}"
	} else if column.IsUnique {
		return "{constraint: unique}"
	}
	return ""
}

func d2TableRelations(table Table) []string {
	relations := common.NewSet[string]()
	for _, c := range table.Columns {
		for _, ref := range c.ReferTo {
			relations.Put(fmt.Sprintf("%s -> %s", table.Name, ref))
		}
	}
	return relations.Keys()
}

type d2TemplateData struct {
	Direction string
	Groups    []Group
	Tables    []Table
}

var d2TemplateText = `
{{- define "styles"}}
	classes: {
		group: {
			style: {
				border-radius: 20
				stroke-width: 2
				bold: true
			}
		}
		relation: {
			style: {
				stroke-width: 2
			}
		}
	}
{{- end}}

{{- define "d2Tables"}}
	{{template "styles"}}
	direction: {{$.Direction}}
	
	{{- range $tableIdx, $table := .Tables }}
		{{$tableFg := (fgColor $tableIdx)}}
		{{$table.Name}}: {
			shape: sql_table
			style: {
				fill: "{{$tableFg}}"
				stroke: "#FFF"
			}

			{{- range $col := $table.Columns}}
				{{column $col.Name}}: {{$col.Type}} {{constraint $col}}
			{{- end}}
		}
	{{- end}}

	{{- range $tableIdx, $table := .Tables }}
		{{$tableFg := (fgColor $tableIdx)}}
		{{- range $rel := relations $table}}
			{{$rel}}: {class: relation; style.stroke: "{{$tableFg}}"}
		{{- end}}
	{{- end}}

	xxx-bold-node: "|" {
		width: 1
		height: 1
		style: {
			opacity: 0
			stroke-width: 0
			bold: true
		}
	}	  
{{- end}}

{{- define "d2Groups"}}
	{{template "styles"}}
	direction: {{$.Direction}}

	{{- range $groupIdx, $group := .Groups }}
		{{- $groupBg := (bgColor $groupIdx) }}
		{{- $groupFg := (fgColor $groupIdx) }}
		{{$group.Name}}: {{$group.Label}} {
			class: group
			style: {
				fill: "{{$groupBg}}"
				stroke: "{{$groupFg}}"
				font-color: "{{$groupFg}}"
			}

			{{- range $table := $group.Tables}}
				{{$table.Name}}: {
					shape: sql_table
					style: {
						fill: "{{$groupFg}}"
						stroke: "#FFF"
					}

					{{- range $col := $table.Columns}}
						{{column $col.Name}}: {{$col.Type}} {{constraint $col}}
					{{- end}}
				}
			{{- end}}
		}
	{{- end}}

	{{- range $groupIdx, $group := .Groups }}
		{{- $groupFg := (fgColor $groupIdx) }}
		{{- range $table := $group.Tables}}
			{{- range $rel := relations $table}}
				{{$group.Name}}.{{$rel}}: {class: relation; style.stroke: "{{$groupFg}}"}
			{{- end}}
		{{- end}}
	{{- end}}
{{- end}}
`
