package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	flagDemo        = flag.Bool("demo", false, "play automatically? (this cheats to win)")
	flagHelp        = flag.Bool("help", false, "Show this help-text?")
	flagRefreshRate = flag.Int("refresh-rate", 20, "Refresh-rate per second")
	flagSeed        = flag.Int64("seed", 0, "Randomizer Seed value (will use current time if ZERO)")
)

func initFlags() {
	flag.Parse()
	if *flagHelp {
		printHelp()
	}

	if *flagSeed == 0 {
		*flagSeed = time.Now().Unix()
	}
	rand.Seed(*flagSeed)
}

func printHelp() {
	fmt.Println(`go-sudoku: A GoLang implementation of the Sudoku game.

Version: ` + version + `

Flags
=====`)
	flag.PrintDefaults()
	os.Exit(0)
}
