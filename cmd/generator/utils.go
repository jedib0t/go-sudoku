package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

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
