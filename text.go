package prim

import (
	"bytes"
	"strings"

	"github.com/mgutz/ansi"
)

type textMethod func(int)

// Text defines implementation basic things for text
type Text struct {
	output      string
	buffer      bytes.Buffer
	outputColor string
	pipeline    []TextPipleline
}

type TextPipleline struct {
	methods []textMethod
}

func NewText(text string) *Text {
	return &Text{
		output:   text,
		pipeline: []TextPipleline{},
	}
}

// IdentLeft provides idents from left by n symbols
func (t *Text) IdentLeft(n int) *Text {
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
	return t
}

// IdentTop provides ident from top on n symbols
func (t *Text) IdentTop(n int) *Text {
	for i := 0; i < n; i++ {
		t.buffer.WriteString("\n")
	}
	return t
}

func (t *Text) AlignCenter(width int) *Text {
	totalPadLen := width - len(t.output)
	if totalPadLen < 0 {
		totalPadLen = 0
	}
	t.buffer.WriteString(strings.Repeat(" ", totalPadLen/2))
	return t
}

func (t *Text) Color(color string) *Text {
	t.outputColor = ansi.Color(t.output, color)
	return t
}

// Output returns result string
func (t *Text) Output() string {
	return output(t)
}

func output(t *Text) string {
	if len(t.pipeline) == 0 {
		return ""
	}
	if len(t.pipeline) == 1 {

	}

	return t.buffer.String()
}
