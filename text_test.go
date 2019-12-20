package prim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentTop(t *testing.T) {
	output := NewText("output").
		IdentTop(2).
		String()
	assert.Equal(t, output, "\n\noutput")

	output = NewText("output").
		IdentTop(-2).
		String()
	assert.Equal(t, output, "")
}

func TestIdentLeft(t *testing.T) {
	output := NewText("output").
		IdentLeft(2).
		String()
	assert.Equal(t, output, "  output")

	output = NewText("output").
		IdentTop(-2).
		String()
	assert.Equal(t, output, "")
}

func TestRemoveEscapes(t *testing.T) {
	output := NewText("output\x1b").
		RemoveEscapes().
		String()
	assert.Equal(t, output, "output")
}

func TestColor(t *testing.T) {
	output := NewText("output").Color("red").String()
	assert.Equal(t, output, "\033[31moutput\033[0m")

	output = NewText("output").Color("reddddd").String()
	assert.Equal(t, output, "\033[30moutput\033[0m")
}
