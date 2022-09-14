package main

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
