package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSubGrid = SubGrid{
		Locations: []Location{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		grid:      testGridDone,
	}
	testSubGridEmpty   = SubGrid{}
	testSubGridPartial = SubGrid{
		Locations: []Location{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		grid:      testGridPartial,
	}
)

func TestSubGrid_Count(t *testing.T) {
	assert.Equal(t, 9, testSubGrid.Count())
	assert.Equal(t, 6, testSubGridPartial.Count())
}

func TestSubGrid_Done(t *testing.T) {
	assert.Equal(t, true, testSubGrid.Done())
	assert.Equal(t, false, testSubGridPartial.Done())
}

func TestSubGrid_Empty(t *testing.T) {
	assert.Equal(t, true, testSubGridEmpty.Empty())
	assert.Equal(t, false, testSubGrid.Empty())
}

func TestSubGrid_Has(t *testing.T) {
	assert.Equal(t, false, testSubGridPartial.Has(1))
	assert.Equal(t, false, testSubGridPartial.Has(2))
	assert.Equal(t, true, testSubGridPartial.Has(3))
	assert.Equal(t, true, testSubGridPartial.Has(4))
	assert.Equal(t, true, testSubGridPartial.Has(5))
	assert.Equal(t, false, testSubGridPartial.Has(6))
	assert.Equal(t, true, testSubGridPartial.Has(7))
	assert.Equal(t, false, testSubGridPartial.Has(8))
	assert.Equal(t, true, testSubGridPartial.Has(9))
}

func TestSubGrid_HasLocation(t *testing.T) {
	assert.True(t, testSubGrid.HasLocation(0, 0))
	assert.True(t, testSubGrid.HasLocation(0, 1))
	assert.True(t, testSubGrid.HasLocation(0, 2))
	assert.True(t, testSubGrid.HasLocation(1, 0))
	assert.True(t, testSubGrid.HasLocation(1, 1))
	assert.True(t, testSubGrid.HasLocation(1, 2))
	assert.True(t, testSubGrid.HasLocation(2, 0))
	assert.True(t, testSubGrid.HasLocation(2, 1))
	assert.True(t, testSubGrid.HasLocation(2, 2))
	assert.False(t, testSubGrid.HasLocation(3, 3))
}

func TestSubGrid_String(t *testing.T) {
	assert.Equal(t, "[4 7 9 8 2 3 1 6 5]", testSubGrid.String())
}
