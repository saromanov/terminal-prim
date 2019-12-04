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
