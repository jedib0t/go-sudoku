package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku/pattern"
)

var (
	// types of sudoku
	typeDefault = "default"
	typeJigSaw  = "jigsaw"
	typeSamurai = "samurai"

	// possible values
	generators   = strings.ToLower(strings.Join(generator.Generators(), "/"))
	difficulties = strings.ToLower(strings.Join([]string{"none", "easy", "medium", "hard", "insane"}, "/"))
	formats      = strings.ToLower(strings.Join([]string{"csv", "table"}, "/"))
	patterns     = strings.ToLower(strings.Join(pattern.Names(), "/"))
	seed         = int64(0)
	types        = strings.Join([]string{typeDefault, typeJigSaw, typeSamurai}, "/")

	// flags
	flagAlgorithm  = flag.String("algorithm", "back-tracking", "Algorithm ("+generators+")")
	flagDebug      = flag.Bool("debug", false, "Enable Debug Logging?")
	flagDifficulty = flag.String("difficulty", "medium", "Difficulty ("+difficulties+")")
	flagFormat     = flag.String("format", "table", "Rendering Format ("+formats+")")
	flagHelp       = flag.Bool("help", false, "Display this usage and help text")
	flagInput      = flag.String("input", "", "File containing a Sudoku Puzzle in CSV format")
	flagNoColor    = flag.Bool("no-color", false, "Disable colors in rendering? [$NO_COLOR]")
	flagPattern    = flag.String("pattern", "", "Pattern to use instead of Difficulty ("+patterns+")")
	flagProgress   = flag.Bool("progress", false, "Show progress in real-time with an artificial delay?")
	flagSeed       = flag.Int64("seed", 0, "RNG Seed (0 => random number based on time) [$SEED]")
	flagType       = flag.String("type", "default", "Sudoku Type ("+types+")")
)

func initFlags() {
	flag.Parse()
	if *flagHelp {
		printHelp()
		os.Exit(0)
	}

	if *flagNoColor || os.Getenv("NO_COLOR") == "1" || !isInteractiveTerminal() {
		text.DisableColors()
	}
	generator.Seed(getSeed())
}

func printHelp() {
	fmt.Printf(`go-sudoku: A GoLang based Sudoku generator and solver.

Usage: go-sudoku [flags] <action>

Version: ` + version + `

Actions:
--------
  * generate: Generate a Sudoku Grid and apply the specified difficulty on it
  * solve: Solve a Sudoku puzzle provided in a text (CSV) file

Examples:
---------
  * ./go-sudoku
  * ./go-sudoku -algorithm back-tracking -theme green -seed 42 generate
  * ./go-sudoku -format csv generate
  * ./go-sudoku -input puzzle.csv solve
  * ./go-sudoku -difficulty hard -format csv generate | ./go-sudoku solve

Optional Flags:
---------------
`)
	flag.PrintDefaults()
}
