package main

import (
	"flag"
	"strings"
)

var (
	// version
	version = "dev"
)

func main() {
	initFlags()

	switch strings.ToLower(flag.Arg(0)) {
	case "", "generate":
		switch strings.ToLower(*flagType) {
		case typeJigSaw:
			generateJigSawSudoku()
		case typeSamurai:
			generateSamuraiSudoku()
		default:
			generateSudoku()
		}
	case "solve":
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
