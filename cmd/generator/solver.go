package main

import (
	"fmt"
	"math/rand"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
)

func solveSudoku() {
	gen := getGenerator()
	rng := rand.New(rand.NewSource(getSeed()))

	// read input and construct a grid
	grid := sudoku.NewGrid(sudoku.WithRNG(rng))
	err := grid.UnmarshalCSV(readCSV())
	if err != nil {
		logErrorAndExit("failed to parse Sudoku from file '%s': %v\n", *flagInput, err)
	}

	// solve using the generator
	gridSolved, err := gen.Generate(grid)
	if err != nil {
		logErrorAndExit("ERROR: failed to solve Sudoku: %v\n", err)
	}

	// render input and output
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Input", "Output"})
	tw.AppendRow(table.Row{sudoku.Render(grid), sudoku.RenderDiff(gridSolved, grid)})
	tw.AppendFooter(table.Row{
		fmt.Sprintf("%d blocks to solve", grid.CountToDo()),
		fmt.Sprintf("%v cycles", gridSolved.GetMetadata("cycles")),
	})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Input", AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
		{Name: "Output", AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(table.StyleBold)
	tw.Style().Format.Footer = text.FormatDefault
	fmt.Println(tw.Render())

	// verify that no blocks were over-written
	for idx := 0; idx < 81; idx++ {
		row, col := idx/9, idx%9
		ogVal := grid.Get(row, col)
		if ogVal != 0 {
			val := gridSolved.Get(row, col)
			if ogVal != val {
				logErrorAndExit("input[%d, %d] was %d, but was over-written by solver with %d",
					row+1, col+1, ogVal, val)
			}
		}
	}
}
