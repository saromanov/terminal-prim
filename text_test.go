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
		IdentLeft(2).
		String()
	assert.Equal(t, output, "output")
}
