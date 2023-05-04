package model

type Group struct {
	Name   string
	Label  string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name     string
	Type     string
	IsPK     bool
	IsFK     bool
	IsUnique bool
	ReferTo  []string
}
