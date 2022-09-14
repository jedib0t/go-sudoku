package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
)

var (
	// game state
	cursor     = sudoku.Location{X: 0, Y: 0}
	diff       difficulty.Difficulty
	errorStr   = ""
	grid       *sudoku.Grid // gridAnswer + difficulty applied
	gridAnswer *sudoku.Grid // contains all the locations filled
	gridOG     *sudoku.Grid // (==gridAnswer) for keeping track of user progress
	userQuit   = false

	// utils
	rng *rand.Rand

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
		if grid.Done() || userQuit {
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
	g := generator.BackTrackingGenerator(
		generator.WithRNG(rng),
	)
	var err error
	gridAnswer, err = g.Generate(nil)
	if err != nil {
		panic(err)
	}
	grid = gridAnswer.Clone()
	grid.ApplyDifficulty(diff)
	gridOG = grid.Clone()
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
		if cursor.X+1 < 9 {
			cursor.X++
		}
	case keyboard.KeyArrowUp:
		if cursor.X-1 >= 0 {
			cursor.X--
		}
	case keyboard.KeyArrowRight:
		if cursor.Y+1 < 9 {
			cursor.Y++
		}
	case keyboard.KeyArrowLeft:
		if cursor.Y-1 >= 0 {
			cursor.Y--
		}
	case keyboard.KeyBackspace, keyboard.KeyBackspace2, keyboard.KeyDelete:
		errorStr = ""
		if grid.IsSet(cursor.X, cursor.Y) && !gridOG.IsSet(cursor.X, cursor.Y) {
			grid.Reset(cursor.X, cursor.Y)
		}
	default:
		if char == 'h' || char == 'H' {
			handleActionHelp()
		} else if char == 'q' || char == 'Q' {
			handleActionQuit()
		} else if char >= '1' && char <= '9' {
			charNum, _ := strconv.Atoi(string(char))
			if charNum >= 1 && charNum <= 9 {
				errorStr = ""
				if !grid.IsSet(cursor.X, cursor.Y) {
					if !*flagAllowWrong && gridAnswer.Get(cursor.X, cursor.Y) != charNum {
						errorStr = fmt.Sprintf("%d is not right @(%d, %d)", charNum, cursor.X+1, cursor.Y+1)
					} else if !grid.Set(cursor.X, cursor.Y, charNum) {
						errorStr = fmt.Sprintf("%d does not fit @(%d, %d)", charNum, cursor.X+1, cursor.Y+1)
					}
				}
			}
		} else {
			handleActionInput(char)
		}
	}
}
