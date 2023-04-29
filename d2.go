package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func generateD2codes(groups []Group) (string, error) {
	if len(groups) == 1 {
		return genD2FromTables(groups[0].Tables)
	} else {
		return genD2FromGroups(groups)
	}
}

func genD2FromTables(tables []Table) (string, error) {
	tpl, err := template.New("d2-code").Parse(d2TablesTemplate)
	if err != nil {
		return "", fmt.Errorf("template error: %w", err)
	}

	buffer := bytes.NewBuffer(nil)
	if err = tpl.Execute(buffer, tables); err != nil {
		return "", fmt.Errorf("template write error: %w", err)
	}

	return buffer.String(), nil
}

func genD2FromGroups(groups []Group) (string, error) {
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
	tpl, err := template.New("d2-code").Parse(d2GroupsTemplate)
	if err != nil {
		return "", fmt.Errorf("template error: %w", err)
	}

	buffer := bytes.NewBuffer(nil)
	if err = tpl.Execute(buffer, groups); err != nil {
		return "", fmt.Errorf("template write error: %w", err)
	}

	return buffer.String(), nil
}

var d2TablesTemplate = `direction: right
{{- range $table := . }}
{{$table.Name}}: {
	shape: sql_table

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

{{- range $table := .}}
{{- range $rel := $table.RelatedTables}}
{{$table.Name}} -> {{$rel}}
{{- end}}
{{- end}}
`

var d2GroupsTemplate = `direction: right
{{- range $group := .}}
{{$group.Name}}: {{$group.Label}} {
	{{- range $table := $group.Tables}}
	{{$table.Name}}: {
		shape: sql_table

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

{{- range $group := .}}
{{- range $table := $group.Tables}}
{{- range $rel := $table.RelatedTables}}
{{$group.Name}}.{{$table.Name}} -> {{$rel}}
{{- end}}
{{- end}}
{{- end}}
`
