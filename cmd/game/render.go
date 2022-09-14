package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
)

var (
	// colors
	colorNumbers = map[string][3]text.Colors{ // bg1 color, bg2 color, letter color
		"answer":   {{text.FgBlack}, {text.BgBlack}, {text.FgHiCyan, text.Italic}},
		"bad":      {{text.FgBlack}, {text.BgBlack}, {text.FgHiRed}},
		"cursor":   {{text.FgBlue}, {text.BgBlue}, {text.FgBlack}},
		"key":      {{text.FgBlack}, {text.BgBlack}, {text.FgHiWhite}},
		"key.done": {{text.FgBlack}, {text.BgBlack}, {text.FgHiBlack}},
		"key.hint": {{text.FgBlack}, {text.BgBlack}, {text.FgHiCyan, text.Italic}},
		"og":       {{text.FgBlack}, {text.BgBlack}, {text.FgWhite}},
		"selected": {{text.FgGreen}, {text.BgGreen}, {text.FgHiWhite}},
	}
	styleDefault = table.StyleColoredBlueWhiteOnBlack
	styleSuccess = table.StyleColoredGreenWhiteOnBlack

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

	style := styleDefault
	if grid.Done() {
		style = styleSuccess
	}

	tw := table.NewWriter()
	tw.SetTitle(renderTitle() + "\n")
	tw.AppendHeader(table.Row{renderStats()})
	tw.AppendRow(table.Row{renderGrid()})
	tw.AppendFooter(table.Row{renderKeyboard()})
	tw.AppendFooter(table.Row{renderShortcuts() + "\n"})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(style)
	tw.Style().Options.SeparateRows = false
	tw.Style().Format.Header = text.FormatDefault
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

func renderGrid() string {
	sgLocations := [9]sudoku.Location{
		{X: 0, Y: 0}, {X: 0, Y: 3}, {X: 0, Y: 6},
		{X: 3, Y: 0}, {X: 3, Y: 3}, {X: 3, Y: 6},
		{X: 6, Y: 0}, {X: 6, Y: 3}, {X: 6, Y: 6},
	}

	tw := table.NewWriter()
	twRow := table.Row{}
	valCursor := grid.Get(cursor.X, cursor.Y)
	for _, loc := range sgLocations {
		twSG := table.NewWriter()
		sg := grid.SubGrid(loc.X, loc.Y)

		var row table.Row
		for idx, loc := range sg.Locations {
			val := grid.Get(loc.X, loc.Y)
			valAnswer := gridAnswer.Get(loc.X, loc.Y)
			valOG := gridOG.Get(loc.X, loc.Y)

			colors := colorNumbers["answer"]
			if val == 0 {
				colors = colorNumbers["key"]
			} else if val == valCursor {
				colors = colorNumbers["selected"]
			} else if val == valOG {
				colors = colorNumbers["og"]
			} else if val != valAnswer && *flagShowWrong {
				colors = colorNumbers["bad"]
			}
			if cursor == loc {
				colors[0] = colorNumbers["cursor"][0]
				colors[1] = colorNumbers["cursor"][1]
			}
			row = append(row, renderKey(val, colors))

			if (idx+1)%3 == 0 {
				twSG.AppendRow(row)
				row = table.Row{}
			}
		}
		twSG.Style().Options = table.OptionsNoBordersAndSeparators
		twRow = append(twRow, twSG.Render())

		if len(twRow) == 3 {
			tw.AppendRow(twRow)
			tw.AppendSeparator()
			twRow = table.Row{}
		}
	}
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.DrawBorder = true
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	return tw.Render()
}
