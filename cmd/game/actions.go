package main

import "github.com/jedib0t/go-sudoku/sudoku"

func handleActionHelp() {
	easiestBlock := cursor
	var possibilities *sudoku.Possibilities
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			p := grid.Possibilities(x, y)
			if p != nil && (possibilities == nil || possibilities.AvailableLen() > p.AvailableLen()) {
				easiestBlock = sudoku.Location{X: x, Y: y}
				possibilities = p
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
