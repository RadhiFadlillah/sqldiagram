package main

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func generateD2(groups []Group) (*d2target.Diagram, *d2graph.Graph, error) {
	// Prepare template
	tpl, err := template.New("d2").
		Funcs(template.FuncMap{
			"bgColor": getBgColor,
			"fgColor": getFgColor,
			"column":  d2EscapeKeyword,
		}).Parse(d2TemplateText)
	if err != nil {
		return nil, nil, fmt.Errorf("create template error: %w", err)
	}

	// Execute template to generate D2 codes
	var d2Codes []byte
	if len(groups) == 1 {
		d2Codes, err = d2CodesFromTables(tpl, groups[0].Tables)
	} else {
		d2Codes, err = d2CodesFromGroups(tpl, groups)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("exec template error: %w", err)
	}

	// Prepare text measurer
	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return nil, nil, fmt.Errorf("textmeasure error: %w", err)
	}

	// Prepare ELK layout
	layout := func(ctx context.Context, g *d2graph.Graph) error {
		return d2elklayout.Layout(ctx, g, &d2elklayout.ConfigurableOpts{
			Algorithm:       "layered",
			NodeSpacing:     20.0,
			Padding:         "[top=50,left=50,bottom=50,right=50]",
			EdgeNodeSpacing: 50.0,
			SelfLoopSpacing: 50.0,
		})
	}

	// Generate diagram and graph
	ctx := context.Background()
	diagram, graph, err := d2lib.Compile(ctx, string(d2Codes), &d2lib.CompileOptions{
		Ruler:  ruler,
		Layout: layout})
	if err != nil {
		return nil, nil, fmt.Errorf("compile error: %w", err)
	}

	return diagram, graph, nil
}

func d2CodesFromTables(tpl *template.Template, tables []Table) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := tpl.ExecuteTemplate(buffer, "d2Tables", tables); err != nil {
		return nil, fmt.Errorf("template write error: %w", err)
	}

	return buffer.Bytes(), nil
}

func d2CodesFromGroups(tpl *template.Template, groups []Group) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := tpl.ExecuteTemplate(buffer, "d2Groups", groups); err != nil {
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
	direction: right
	
	{{- range $tableIdx, $table := . }}
		{{$tableFg := (fgColor $tableIdx)}}
		{{$table.Name}}: {
			shape: sql_table
			style: {
				fill: "{{$tableFg}}"
				stroke: "#FFF"
			}

			{{- range $pk := $table.PrimaryKeys}}
				{{column $pk.Name}}: {{$pk.Tp}} {constraint: primary_key}
			{{- end}}

			{{- range $fk := $table.ForeignKeys}}
				{{column $fk.Name}}: {{$fk.Tp}} {constraint: foreign_key}
			{{- end}}

			{{- range $col := $table.Columns}}
				{{column $col.Name}}: {{$col.Tp}} {{- if $col.Unique -}} {constraint: unique} {{- end}}
			{{- end}}
		}
	{{- end}}

	{{- range $tableIdx, $table := . }}
		{{$tableFg := (fgColor $tableIdx)}}
		{{- range $rel := $table.RelatedTables}}
			{{$table.Name}} -> {{$rel}}: {class: relation; style.stroke: "{{$tableFg}}"}
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
	direction: right

	{{- range $groupIdx, $group := .}}
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

					{{- range $pk := $table.PrimaryKeys}}
						{{column $pk.Name}}: {{$pk.Tp}} {constraint: primary_key}
					{{- end}}

					{{- range $fk := $table.ForeignKeys}}
						{{column $fk.Name}}: {{$fk.Tp}} {constraint: foreign_key}
					{{- end}}

					{{- range $col := $table.Columns}}
						{{column $col.Name}}: {{$col.Tp}} {{- if $col.Unique -}} {constraint: unique} {{- end}}
					{{- end}}
				}
			{{- end}}
		}
	{{- end}}

	{{- range $groupIdx, $group := .}}
		{{- $groupFg := (fgColor $groupIdx) }}
		{{- range $table := $group.Tables}}
			{{- range $rel := $table.RelatedTables}}
				{{$group.Name}}.{{$table.Name}} -> {{$rel}}: {class: relation; style.stroke: "{{$groupFg}}"}
			{{- end}}
		{{- end}}
	{{- end}}
{{- end}}
`
