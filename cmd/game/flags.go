package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
)

var (
	difficulties = strings.ToLower(strings.Join([]string{"none", "easy", "medium", "hard", "insane"}, "/"))

	flagAllowWrong  = flag.Bool("allow-wrong", false, "Allow incorrect values?")
	flagDemo        = flag.Bool("demo", false, "play automatically? (this cheats to win)")
	flagDifficulty  = flag.String("difficulty", "medium", "Difficulty ("+difficulties+")")
	flagHelp        = flag.Bool("help", false, "Show this help-text?")
	flagHints       = flag.Bool("hints", false, "Highlight possible values in the Keyboard?")
	flagRefreshRate = flag.Int("refresh-rate", 20, "Refresh-rate per second")
	flagSeed        = flag.Int64("seed", 0, "Randomizer Seed value (will use current time if ZERO)")
	flagShowWrong   = flag.Bool("show-wrong", false, "Highlight incorrect values in Red?")
)

func initFlags() {
	flag.Parse()
	if *flagHelp {
		printHelp()
	}

	diff = difficulty.From(*flagDifficulty)
	if *flagSeed == 0 {
		*flagSeed = time.Now().UnixNano()
	}
	rng = rand.New(rand.NewSource(*flagSeed))
}

func printHelp() {
	fmt.Println(`go-sudoku: A GoLang implementation of the Sudoku game.

Version: ` + version + `

Flags
=====`)
	flag.PrintDefaults()
	os.Exit(0)
}
