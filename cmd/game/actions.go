package main

import "github.com/jedib0t/go-sudoku/sudoku"

func handleActionHelp() {
	numOccurringMost, numOccurrencesMost := 1, 0
	for n := 1; n <= 9; n++ {
		numOccurrences := grid.CountValue(n)
		if numOccurrences > numOccurrencesMost {
			numOccurringMost = n
			numOccurrencesMost = numOccurrences
		}
	}

	easiestBlock := cursor
	var possibilities *sudoku.Possibilities
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			p := grid.Possibilities(x, y)
			if p != nil && (possibilities == nil || possibilities.AvailableLen() > p.AvailableLen()) {
				easiestBlock = sudoku.Location{X: x, Y: y}
				possibilities = p

				// stop early if we've found a number that occurs most, and fits
				// a block like nothing else
				if possibilities.AvailableLen() == 1 && possibilities.AvailableMap()[numOccurringMost] {
					cursor = easiestBlock
					return
				}
			}
		}
	}
	cursor = easiestBlock
}

func handleActionQuit() {
	userQuit = true
}

func handleActionReset() {
	renderMutex.Lock()
	defer renderMutex.Unlock()

	generateSudoku()
}

func handleActionInput(char rune) {
	renderMutex.Lock()
	defer renderMutex.Unlock()

}
