package prim

import (
	"bytes"
	"strings"
)

// Text defines implementation basic things for text
type Text struct {
	output string
	buffer bytes.Buffer
}

func NewText(text string) *Text {
	return &Text{
		output: text,
	}
}

// IdentLeft provides idents from left by n symbols
func (t *Text) IdentLeft(n int) {
	result := strings.Repeat(" ", n)
	lines := strings.Split(t.output, "\n")
	for _, line := range lines {
		t.buffer.WriteString(result)
		t.buffer.WriteString(line)
	}
	t.output = result
}

// IdentTop provides ident from top on n symbols
func (t *Text) IdentTop(n int) {
	for i := 0; i < n; i++ {
		t.buffer.WriteString("\n")
	}
	t.buffer.WriteString(t.output)
}

// Output returns result string
func (t *Text) Output() string {
	return t.buffer.String()
}
