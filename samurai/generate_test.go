package samurai

import (
	"testing"

	"github.com/jedib0t/go-sudoku/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	gen := generator.BackTrackingGenerator()
	grids, err := Generate(gen)
	assert.Len(t, grids, 5)
	assert.Nil(t, err)
	assert.Equal(t, 81, grids[0].Count())
	assert.Equal(t, 81, grids[1].Count())
	assert.Equal(t, 81, grids[2].Count())
	assert.Equal(t, 81, grids[3].Count())
	assert.Equal(t, 81, grids[4].Count())
	assert.Equal(t, grids[0].SubGrid(0, 0).String(), grids[1].SubGrid(6, 6).String())
	assert.Equal(t, grids[0].SubGrid(0, 6).String(), grids[2].SubGrid(6, 0).String())
	assert.Equal(t, grids[0].SubGrid(6, 0).String(), grids[3].SubGrid(0, 6).String())
	assert.Equal(t, grids[0].SubGrid(6, 6).String(), grids[4].SubGrid(0, 0).String())
}
