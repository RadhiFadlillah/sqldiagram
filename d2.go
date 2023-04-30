package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func generateD2codes(groups []Group) (string, error) {
	tpl, err := template.New("d2").
		Funcs(template.FuncMap{"bgColor": getBgColor}).
		Funcs(template.FuncMap{"fgColor": getFgColor}).
		Parse(d2TemplateText)
	if err != nil {
		return "", fmt.Errorf("template error: %w", err)
	}

	if len(groups) == 1 {
		return genD2FromTables(tpl, groups[0].Tables)
	} else {
		return genD2FromGroups(tpl, groups)
	}
}

func genD2FromTables(tpl *template.Template, tables []Table) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if err := tpl.ExecuteTemplate(buffer, "d2Tables", tables); err != nil {
		return "", fmt.Errorf("template write error: %w", err)
	}

	return buffer.String(), nil
}

func genD2FromGroups(tpl *template.Template, groups []Group) (string, error) {
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
			for i, rt := range t.RelatedTables {
				if rtGroup, exist := mapTableGroup[rt]; exist {
					t.RelatedTables[i] = rtGroup + "." + rt
				}
			}
		}
	}

	// Execute template
	buffer := bytes.NewBuffer(nil)
	if err := tpl.ExecuteTemplate(buffer, "d2Groups", groups); err != nil {
		return "", fmt.Errorf("template write error: %w", err)
	}

	return buffer.String(), nil
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
				{{$pk.Name}}: {{$pk.Tp}} {constraint: primary_key}
			{{- end}}

			{{- range $fk := $table.ForeignKeys}}
				{{$fk.Name}}: {{$fk.Tp}} {constraint: foreign_key}
			{{- end}}

			{{- range $col := $table.Columns}}
				{{$col.Name}}: {{$col.Tp}} {{- if $col.Unique -}} {constraint: unique} {{- end}}
			{{- end}}
		}
	{{- end}}

	{{- range $tableIdx, $table := . }}
		{{$tableFg := (fgColor $tableIdx)}}
		{{- range $rel := $table.RelatedTables}}
			{{$table.Name}} -> {{$rel}}: {class: relation; style.stroke: "{{$tableFg}}"}
		{{- end}}
	{{- end}}
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
						{{$pk.Name}}: {{$pk.Tp}} {constraint: primary_key}
					{{- end}}

					{{- range $fk := $table.ForeignKeys}}
						{{$fk.Name}}: {{$fk.Tp}} {constraint: foreign_key}
					{{- end}}

					{{- range $col := $table.Columns}}
						{{$col.Name}}: {{$col.Tp}} {{- if $col.Unique -}} {constraint: unique} {{- end}}
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
