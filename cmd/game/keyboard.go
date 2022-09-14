package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func cleanupKeyboard() {
	_ = keyboard.Close()
	cursorShow()
}

func cursorHide() {
	fmt.Printf("\x1b[?25l")
}

func cursorShow() {
	fmt.Printf("\x1b[?25h")
}

func initKeyboard() {
	// over-ride keyboard handling
	if err := keyboard.Open(); err != nil {
		logErrorAndExit(err.Error())
	}
	cursorHide()

	// ensure cleanupKeyboard gets called on exit
	exitHandlers = append(exitHandlers, cleanupKeyboard)
}
