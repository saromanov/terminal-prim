package prim

import (
	"github.com/saromanov/tables"
)

// Table provides representation of the table
type Table struct {
	table *tables.App
}

func NewTable() *Table {
	return &Table{
		table: tables.New(),
	}
}

func (t *Table) AddHeaders(headers ...interface{}) {
	t.table.AddHeader(headers...)
}

func (t *Table) AddLines(lines ...interface{}) {
	t.table.AddLine(lines...)
}

func (t *Table) Output() {
	t.table.Build()
}
