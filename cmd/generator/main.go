package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
	"github.com/jedib0t/go-sudoku/sudoku/pattern"
	"github.com/jedib0t/go-sudoku/variants/jigsaw"
	"github.com/jedib0t/go-sudoku/variants/samurai"
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

	// state
	showProgressLinesShown = 0
	showProgressInterval   = time.Millisecond * 25

	// version
	version = "dev"
)

func main() {
	flag.Parse()
	if *flagHelp {
		printHelp()
		return
	}

	switch strings.ToLower(flag.Arg(0)) {
	case "", "generate":
		initStuff()
		switch strings.ToLower(*flagType) {
		case typeJigSaw:
			generateJigSawSudoku()
		case typeSamurai:
			generateSamuraiSudoku()
		default:
			generateSudoku()
		}
	case "solve":
		initStuff()
		switch strings.ToLower(*flagType) {
		case typeDefault:
			solveSudoku()
		default:
			logErrorAndExit("logic to solve %s sudoku unimplemented", *flagType)
		}
	default:
		printHelp()
	}
}

func applyPatternOrDifficulty(grids ...*sudoku.Grid) string {
	diff := getDifficulty()
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

func getDifficulty() difficulty.Difficulty {
	switch strings.ToLower(*flagDifficulty) {
	case "none":
		return difficulty.None
	case "easy":
		return difficulty.Easy
	case "medium":
		return difficulty.Medium
	case "hard":
		return difficulty.Hard
	case "insane":
		return difficulty.Insane
	default:
		return difficulty.Medium
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

func getSeed() int64 {
	if seed == 0 {
		seed = *flagSeed
	}
	if seed == 0 {
		seedStr := os.Getenv("SEED")
		seedVal, err := strconv.ParseInt(seedStr, 10, 64)
		if err == nil {
			seed = seedVal
		}
	}
	if seed == 0 {
		seed = (time.Now().UnixNano() / 1000) % 10000
	}
	return seed
}

func initStuff() {
	if *flagNoColor || os.Getenv("NO_COLOR") == "1" || !isInteractiveTerminal() {
		text.DisableColors()
	}
	generator.Seed(getSeed())
}

func isInteractiveTerminal() bool {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}

func logErrorAndExit(msg string, a ...interface{}) {
	errMsg := text.FgHiRed.Sprintf("ERROR: "+msg, a...)
	_, _ = fmt.Fprintln(os.Stderr, errMsg)
	os.Exit(1)
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

func readCSV() string {
	var csvString string
	if *flagInput == "" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logErrorAndExit("failed to read Sudoku from stdin: %v\n", err)
		}
		csvString = string(bytes)
	} else {
		bytes, err := ioutil.ReadFile(*flagInput)
		if err != nil {
			logErrorAndExit("failed to read Sudoku from file '%s': %v\n", *flagInput, err)
		}
		csvString = string(bytes)
	}
	return csvString
}

func renderGeneratedResult(title string, result string) {
	tw := table.NewWriter()
	tw.AppendRow(table.Row{result})
	tw.AppendFooter(table.Row{fmt.Sprintf("%s [%v]", title, getSeed())})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(table.StyleBold)
	tw.Style().Format.Footer = text.FormatDefault
	fmt.Println(tw.Render())
}

func renderProgress(g, og *sudoku.Grid, attempts, cycles int) {
	for showProgressLinesShown > 0 {
		fmt.Print(text.CursorUp.Sprint())
		fmt.Print(text.EraseLine.Sprint())
		showProgressLinesShown--
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleBold)
	switch g.GetMetadata("type") {
	case "jigsaw":
		if og == nil {
			tw.AppendRow(table.Row{jigsaw.Render(g)})
		} else {
			tw.AppendRow(table.Row{jigsaw.RenderDiff(g, og)})
		}
	default:
		if og == nil {
			tw.AppendRow(table.Row{sudoku.Render(g)})
		} else {
			tw.AppendRow(table.Row{sudoku.RenderDiff(g, og)})
		}
	}
	tw.AppendFooter(table.Row{fmt.Sprintf("Attempt # %d.%d", attempts, cycles)})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.Style().Format.Footer = text.FormatDefault
	out := tw.Render()

	fmt.Println(out)
	showProgressLinesShown = strings.Count(out, "\n") + 1
	time.Sleep(showProgressInterval)
}

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
