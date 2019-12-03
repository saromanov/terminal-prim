package prim

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
)

type textMethod func()

// Text defines implementation basic things for text
type Text struct {
	output      string
	buffer      bytes.Buffer
	outputColor string
	pipeline    TextPipleline
}

type TextPipleline struct {
	methods []textMethod
}

func NewText(text string) *Text {
	return &Text{
		output:   text,
		pipeline: TextPipleline{},
	}
}

// IdentLeft provides idents from left by n symbols
func (t *Text) IdentLeft(n int) *Text {
	t.pipeline.methods = append(t.pipeline.methods, func() {
		t.buffer.WriteString(strings.Repeat(" ", n))
	})
	return t
}

// IdentTop provides ident from top on n symbols
func (t *Text) IdentTop(n int) *Text {
	t.pipeline.methods = append(t.pipeline.methods, func() {
		for i := 0; i < n; i++ {
			t.buffer.WriteString("\n")
		}
	})
	return t
}

// Text overwrites current main string
func (t *Text) Text(str string) *Text {
	t.output = str
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
	t.pipeline.methods = append(t.pipeline.methods, func() {
		t.outputColor = ansi.Color(t.output, color)
	})
	return t
}

// Output returns result string
func (t *Text) Output() *Text {
	fmt.Println(output(t))
	t.output = ""
	t.outputColor = ""
	t.pipeline = TextPipleline{}
	return t
}

func output(t *Text) string {
	if len(t.pipeline.methods) == 0 {
		return ""
	}
	for _, pipe := range t.pipeline.methods {
		pipe()
	}
	if t.outputColor != "" {
		t.buffer.WriteString(t.outputColor)
	} else {
		t.buffer.WriteString(t.output)
	}
	return t.buffer.String()
}
