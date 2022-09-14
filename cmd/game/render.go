package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	// colors
	colorsFlaggedWrong = text.Colors{text.FgHiRed}

	// misc
	linesRendered = 0
	renderedGame  = ""

	// controls
	renderEnabled = true
	renderMutex   = sync.Mutex{}
)

func renderAsync(chStop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.Tick(time.Second / time.Duration(*flagRefreshRate))
	for {
		select {
		case <-chStop: // render one final time and return
			renderGame()
			return
		case <-timer: // render as part of regular cycle
			renderGame()
		}
	}
}

func renderGame() {
	renderMutex.Lock()
	defer renderMutex.Unlock() // unnecessary
	if !renderEnabled {
		return
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{renderTitle()})
	tw.AppendHeader(table.Row{renderHeader()})
	tw.AppendRow(table.Row{renderSudoku()})
	tw.AppendFooter(table.Row{renderFooter()})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(table.StyleBold)
	tw.Style().Options.SeparateRows = true
	tw.Style().Format.Footer = text.FormatDefault
	tw.Style().Title.Align = text.AlignCenter

	out := tw.Render()
	if out != renderedGame {
		for linesRendered > 0 {
			fmt.Print(text.CursorUp.Sprint())
			fmt.Print(text.EraseLine.Sprint())
			linesRendered--
		}

		linesRendered = strings.Count(out, "\n") + 1
		fmt.Println(out)
		renderedGame = out
	}
}

func renderSudoku() string {
	return "sudoku game grid goes here and will\nbe the most amazing thing ever"
}
