package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock_Value(t *testing.T) {
	b := Block{value: 13}
	assert.Equal(t, 13, b.Value())
}
