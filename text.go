package prim

import (
	"bytes"
	"strings"

	"github.com/mgutz/ansi"
)

// Text defines implementation basic things for text
type Text struct {
	output      string
	buffer      bytes.Buffer
	outputColor string
}

func NewText(text string) *Text {
	return &Text{
		output: text,
	}
}

// IdentLeft provides idents from left by n symbols
func (t *Text) IdentLeft(n int) {
	buffer := t.buffer
	result := strings.Repeat(" ", n)
	lines := strings.Split(t.output, "\n")
	for _, line := range lines {
		buffer.WriteString(result)
		if t.outputColor != "" {
			buffer.WriteString(t.outputColor)
			continue
		}
		buffer.WriteString(line)
	}
	t.buffer = buffer
	t.output = result
}

// IdentTop provides ident from top on n symbols
func (t *Text) IdentTop(n int) {
	for i := 0; i < n; i++ {
		t.buffer.WriteString("\n")
	}
}

func (t *Text) AlignCenter(width int) {
	totalPadLen := width - len(t.output)
	if totalPadLen < 0 {
		totalPadLen = 0
	}
	t.buffer.WriteString(strings.Repeat(" ", totalPadLen/2))
}

func (t *Text) Color(color string) {
	t.outputColor = ansi.Color(t.output, color)
}

// Output returns result string
func (t *Text) Output() string {
	return t.buffer.String()
}
