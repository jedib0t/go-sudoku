package jigsaw

import (
	"math/rand"

	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku"
)

// Generate generates a JigSaw Sudoku grid.
func Generate(g generator.Generator) (*sudoku.Grid, error) {
	// generate a new grid with the sub-grids
	grid := sudoku.NewGrid(sudoku.WithSubGrids(randomizeSubGrids(g.RNG())))
	grid.SetMetadata("type", "jigsaw")

	// run the generator to fill the numbers
	return g.Generate(grid)
}

func randomizeSubGrids(rng *rand.Rand) []sudoku.SubGrid {
	ptrn := patterns[rng.Intn(len(patterns))]

	subGrids := make([]sudoku.SubGrid, 9)
	for row := range ptrn {
		for col := range ptrn[row] {
			sgIdx := ptrn[row][col] - 1
			loc := sudoku.Location{X: row, Y: col}
			subGrids[sgIdx].Locations = append(subGrids[sgIdx].Locations, loc)
		}
	}

	return subGrids
}
