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

func (t *Text) Ident(n int) {
	result := strings.Repeat(" ", n)
	lines := strings.Split(t.output, "\n")
	for _, line := range lines {
		t.buffer.WriteString(result)
		t.buffer.WriteString(line)
	}
	t.output = result
}
