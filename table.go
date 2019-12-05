package prim

import (
	"github.com/saromanov/tables"
)

// Table provides representation of the table
type Table struct {
	table *tables.App
}

// NewTable provides creating of the Table representation
func NewTable() *Table {
	return &Table{
		table: tables.New(),
	}
}

// AddHeaders provides adding of headers to the table
func (t *Table) AddHeaders(headers ...interface{}) {
	t.table.AddHeader(headers...)
}

// AddLines provides adding of the lines to the table
func (t *Table) AddLines(lines ...interface{}) {
	t.table.AddLine(lines...)
}

// Output provides output table to console
// Example:
// tab := term.NewTable()
// tab.AddHeaders("aaa", "vvvv", "cccc")
// tab.AddLines("1", "2", "3")
// tab.Output()
//
// aaa  vvvv  cccc
// ---  ----  ----
// 1    2     3
func (t *Table) Output() {
	t.table.Build()
}
