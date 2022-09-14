package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// exitHandlers contains all functions that need to be called during exit
	exitHandlers []func()
	// timeStart is used to render the game timer
	timeStart = time.Now()
	// version
	version = "dev"
)

func init() {
	// seed the RNG or the word auto-selected will remain the same all the time
	rand.Seed(time.Now().UnixNano())

	// cleanup on termination
	c := make(chan os.Signal, 5)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c
		cleanup()
	}()

	// init other things
	initFlags()
	initKeyboard()
	initHeaderAndFooter()
}

func cleanup() {
	renderMutex.Lock()
	defer renderMutex.Unlock() // unnecessary
	renderEnabled = false

	for _, exitHandler := range exitHandlers {
		exitHandler()
	}
}

func main() {
	defer cleanup()

	generateSudoku()

	play()
}
