package jigsaw

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/jedib0t/go-sudoku/generator"
	"github.com/stretchr/testify/assert"
)

var (
	testRNG = rand.New(rand.NewSource(42))
)

func TestGenerate(t *testing.T) {
	gen := generator.BackTrackingGenerator(generator.WithRNG(testRNG))

	grid, err := Generate(gen)
	assert.NotNil(t, grid)
	assert.Nil(t, err)
	expectedSubGrids := `1 1 2 2 2 2 2 3 3
1 1 1 2 2 2 3 3 3
4 1 1 1 2 3 3 3 6
4 4 1 5 5 5 3 6 6
4 4 4 5 5 5 6 6 6
4 4 7 5 5 5 9 6 6
4 7 7 7 8 9 9 9 6
7 7 7 8 8 8 9 9 9
7 7 8 8 8 8 8 9 9`
	actualSubGrids := grid.SubGrids().String()
	assert.Equal(t, expectedSubGrids, actualSubGrids)
	if expectedSubGrids != actualSubGrids {
		fmt.Println(actualSubGrids)
	}
}
