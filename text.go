package prim

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/mgutz/ansi"
)

type textMethod func()

// Text defines implementation basic things for text
type Text struct {
	output      string
	buffer      bytes.Buffer
	outputColor string
	pipeline    TextPipleline
	textLines   uint
	methods     map[uint][]textMethod
	lines       map[uint]string
}

type TextPipleline struct {
	methods []textMethod
}

func NewText(text string) *Text {
	lines := map[uint]string{0: text}
	return &Text{
		output:  text,
		methods: map[uint][]textMethod{},
		lines:   lines,
	}
}

// IdentLeft provides idents from left by n symbols
func (t *Text) IdentLeft(n int) *Text {
	if n < 0 {
		return t
	}
	t.addMethod(func() {
		t.buffer.WriteString(strings.Repeat(" ", n))
	})
	return t
}

// IdentTop provides ident from top on n symbols
func (t *Text) IdentTop(n int) *Text {
	if n < 0 {
		return t
	}
	t.addMethod(func() {
		for i := 0; i < n; i++ {
			t.buffer.WriteString("\n")
		}
	})
	return t
}

// Text overwrites current main string
func (t *Text) Text(str string) *Text {
	t.output = str
	t.textLines++
	t.lines[t.textLines] = str
	return t
}

// AlignCenter provides adding text to the center
func (t *Text) AlignCenter(width int) *Text {
	t.addMethod(func() {
		totalPadLen := width - len(t.output)
		if totalPadLen < 0 {
			totalPadLen = 0
		}
		t.buffer.WriteString(strings.Repeat(" ", totalPadLen/2))
	})
	return t
}

// Color provides coloring of the text
func (t *Text) Color(color string) *Text {
	t.addMethod(func() {
		t.lines[t.textLines] = ansi.Color(t.output, color)
		t.outputColor = ansi.Color(t.output, color)
	})
	return t
}

// TrimSpace provides removing of the spaces
func (t *Text) TrimSpace() *Text {
	left := 0
	output := strings.TrimLeftFunc(t.output, func(r rune) bool {
		if unicode.IsSpace(r) {
			left++
			return true
		}
		return false
	})

	output = strings.TrimRightFunc(output, unicode.IsSpace)
	t.lines[t.textLines] = output
	return t
}

// RemoveEscapes provides removing of escapes from the line
func (t *Text) RemoveEscapes() *Text {
	t.addMethod(func() {
		var output []rune
		for _, line := range t.lines {
			for _, r := range []rune(line) {
				if r == '\x1b' {
					continue
				}
				output = append(output, r)
			}
		}
		t.lines[t.textLines] = string(output)
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

// String retruns string representation of the text
func (t *Text) String() string {
	return output(t)
}

func (t *Text) addMethod(method func()) {
	t.pipeline.methods = append(t.pipeline.methods, method)
	t.methods[t.textLines] = append(t.methods[t.textLines], method)
}

func output(t *Text) string {
	if len(t.pipeline.methods) == 0 {
		return ""
	}
	var i uint
	for {
		if i > t.textLines {
			break
		}
		pipes, _ := t.methods[i]
		for _, pipe := range pipes {
			pipe()
		}
		t.buffer.WriteString(t.lines[i])
		i++

	}
	return t.buffer.String()
}
