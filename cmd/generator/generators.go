package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
	"github.com/jedib0t/go-sudoku/sudoku/pattern"
	"github.com/jedib0t/go-sudoku/variants/jigsaw"
	"github.com/jedib0t/go-sudoku/variants/samurai"
)

func applyPatternOrDifficulty(grids ...*sudoku.Grid) string {
	diff := difficulty.From(*flagDifficulty)
	ptrn := pattern.Get(*flagPattern)

	// apply pattern or diff
	if ptrn != nil {
		for _, grid := range grids {
			grid.ApplyPattern(ptrn)
		}
	} else {
		for _, grid := range grids {
			grid.ApplyDifficulty(diff)
		}
	}

	// return the mode
	mode := diff.String()
	if *flagPattern != "" {
		mode = text.FormatTitle.Apply(*flagPattern)
	}
	return mode
}

func generateSudoku() {
	gen := getGenerator()

	// generate the puzzle
	grid, err := gen.Generate(nil)
	if err != nil {
		logErrorAndExit("failed to generate sudoku with %s method: %v", gen.Name(), err)
	}
	// apply pattern or difficulty
	mode := applyPatternOrDifficulty(grid)

	// render
	switch strings.ToLower(*flagFormat) {
	case "csv":
		fmt.Println(grid)
	default:
		renderGeneratedResult(mode, sudoku.Render(grid))
	}
}

func generateJigSawSudoku() {
	gen := getGenerator()

	// generate the puzzle
	grid, err := jigsaw.Generate(gen)
	if err != nil {
		logErrorAndExit("failed to generate jigsaw sudoku with %s method: %v", gen.Name(), err)
	}
	// apply pattern or difficulty
	mode := applyPatternOrDifficulty(grid)

	// render
	switch strings.ToLower(*flagFormat) {
	case "csv":
		fmt.Println(grid)
	default:
		renderGeneratedResult(fmt.Sprintf("JigSaw %s", mode), jigsaw.Render(grid))
	}
}

func generateSamuraiSudoku() {
	gen := getGenerator()

	// generate the puzzle
	grids, err := samurai.Generate(gen)
	if err != nil {
		logErrorAndExit("failed to generate samurai sudoku with %s method: %v", gen.Name(), err)
	}
	if *flagDebug {
		tw := table.NewWriter()
		tw.AppendRow(table.Row{sudoku.Render(grids[1]), " ", sudoku.Render(grids[2])})
		tw.AppendRow(table.Row{" ", sudoku.Render(grids[0]), " "})
		tw.AppendRow(table.Row{sudoku.Render(grids[3]), " ", sudoku.Render(grids[4])})
		tw.SetStyle(table.StyleLight)
		tw.Style().Options.SeparateRows = true
		_, _ = fmt.Fprintln(os.Stderr, tw.Render())
	}
	// apply pattern or difficulty
	mode := applyPatternOrDifficulty(grids...)

	// render
	samuraiSudoku, err := samurai.Render(grids)
	if err != nil {
		logErrorAndExit("failed to generate sudoku with %s method: %v", gen.Name(), err)
	}
	switch strings.ToLower(*flagFormat) {
	case "csv":
		logErrorAndExit("cannot render a Samurai Sudoku in CSV format (unsupported)")
	default:
		renderGeneratedResult(fmt.Sprintf("Samurai %s", mode), samuraiSudoku)
	}
}

func getGenerator() generator.Generator {
	var opts []generator.Option
	if *flagDebug {
		opts = append(opts, generator.WithDebug())
	}
	if *flagProgress && isInteractiveTerminal() {
		opts = append(opts, generator.WithProgress(renderProgress))
	}

	switch strings.ToLower(*flagAlgorithm) {
	case "brute-force":
		return generator.BruteForceGenerator(opts...)
	case "back-tracking":
		return generator.BackTrackingGenerator(opts...)
	default:
		return generator.BackTrackingGenerator(opts...)
	}
}
