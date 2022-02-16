package generator

import (
	"math/rand"
	"testing"

	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
	"github.com/stretchr/testify/assert"
)

func TestBackTrackingGenerator_Generate(t *testing.T) {
	rng := rand.New(rand.NewSource(13))

	// generate a sudoku grid
	grid, err := BackTrackingGenerator(WithRNG(rng)).Generate(nil)
	assert.NotNil(t, grid)
	assert.Nil(t, err)
	gridCSV := grid.MarshalCSV()
	expectedGridCSV := `8,1,5,9,2,7,4,6,3
4,7,6,1,8,3,5,9,2
3,2,9,4,5,6,8,1,7
9,6,7,8,4,1,3,2,5
1,8,2,6,3,5,7,4,9
5,3,4,2,7,9,6,8,1
7,9,8,3,1,4,2,5,6
6,4,3,5,9,2,1,7,8
2,5,1,7,6,8,9,3,4`
	assert.Equal(t, expectedGridCSV, gridCSV)

	// apply a difficulty
	grid.ApplyDifficulty(difficulty.Medium)
	partialGridCSV := grid.MarshalCSV()
	expectedPartialGridCSV := `0,0,0,9,0,7,4,0,3
0,0,0,1,8,0,0,0,0
0,2,0,0,5,0,8,0,7
9,0,7,8,4,1,0,0,5
1,0,0,0,0,0,7,0,9
0,0,0,0,7,9,0,0,1
0,9,8,0,1,0,0,5,0
6,4,0,5,0,2,1,7,8
0,0,1,0,6,0,0,0,4`
	assert.Equal(t, expectedPartialGridCSV, partialGridCSV)

	// solve
	grid, err = BackTrackingGenerator(WithRNG(rng)).Generate(grid)
	assert.NotNil(t, grid)
	assert.Nil(t, err)
	assert.Equal(t, expectedGridCSV, gridCSV)
}

func TestBackTrackingGenerator_Name(t *testing.T) {
	assert.Equal(t, backTrackingGeneratorName, BackTrackingGenerator().Name())
}
