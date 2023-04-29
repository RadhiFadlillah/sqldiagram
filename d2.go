package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func generateD2codes(tables []Table) (string, error) {
	tpl, err := template.New("d2-code").Parse(d2Template)
	if err != nil {
		return "", fmt.Errorf("template error: %w", err)
	}

	buffer := bytes.NewBuffer(nil)
	if err = tpl.Execute(buffer, tables); err != nil {
		return "", fmt.Errorf("template write error: %w", err)
	}

	return buffer.String(), nil
}

var d2Template = `
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
