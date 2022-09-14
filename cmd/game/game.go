package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
)

var (
	// game state
	userQuit = false

	// demo
	demoRNG   = rand.New(rand.NewSource(1))
	demoSpeed = time.Second / 5
)

// play starts the game.
func play() {
	// render forever in a separate routine
	chStop := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go renderAsync(chStop, &wg)

	for {
		if true {
			break
		}

		if *flagDemo {
			demo()
		} else {
			getUserInput()
		}
	}

	renderGame() // one final render
	chStop <- true
	wg.Wait()
}

func demo() {

}

func generateSudoku() {

}

func getUserInput() {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		return
	}
	if *flagDemo && key != keyboard.KeyEsc && key != keyboard.KeyCtrlC && char != 'q' && char != 'Q' {
		return
	}

	switch key {
	case keyboard.KeyEsc, keyboard.KeyCtrlC:
		handleActionQuit()
	case keyboard.KeyCtrlR:
		handleActionReset()
	case keyboard.KeyArrowDown:
		// TODO
	case keyboard.KeyArrowUp:
		// TODO
	case keyboard.KeyArrowRight:
		// TODO
	case keyboard.KeyArrowLeft:
		// TODO
	case keyboard.KeySpace:
		// TODO
	default:
		if char == 'q' || char == 'Q' {
			handleActionQuit()
		} else {
			handleActionInput(char)
		}
	}
}
