package main

import (
	"io/ioutil"
	"os"
)

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
