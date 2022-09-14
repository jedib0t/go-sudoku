package main

import (
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	titleText1 = `
  _________          ___       __          
 /   _____/__ __  __| _/ ____ |  | ____ __ 
 \_____  \|  |  \/ __ | / __ \|  |/ /  |  \
 /        \  |  / /_/ |(  \_\ )    \|  |  /
/_______  /____/\____ | \____/|__|_ \____/ 
        \/           \/            \/      

`
	titleText = titleText1
)

var (
	colorKey   = text.Colors{text.Italic, text.FgWhite}
	colorTitle = []text.Color{
		text.FgHiWhite,
		text.FgHiYellow,
		text.FgHiCyan,
		text.FgHiBlue,
		text.FgHiGreen,
		text.FgHiMagenta,
		text.FgHiRed,
	}
	colorTitleAnimated = false
	colorTitleIdx      = 0
	colorTitleOnce     = sync.Once{}

	titleOnce sync.Once
)

func initHeaderAndFooter() {
	titleOnce.Do(func() {
		// title Colors
		if colorTitleAnimated {
			colorTitleOnce.Do(func() {
				go func() {
					for {
						time.Sleep(time.Second / 2)
						if colorTitleIdx < len(colorTitle)-1 {
							colorTitleIdx++
						} else {
							colorTitleIdx = 0
						}
					}
				}()
			})
		}
	})
}

func renderFooter() string {
	return colorKey.Sprint("navigate: ▶ ▲ ▼ ◀  | quit: <Q/q>")
}

func renderHeader() string {
	tw := table.NewWriter()
	tw.AppendRow(table.Row{"header items go here"})
	tw.Style().Options.DrawBorder = false
	tw.Style().Options.SeparateColumns = false
	tw.Style().Options.SeparateRows = false
	return tw.Render()
}

func renderTitle() string {
	colors := text.Colors{colorTitle[colorTitleIdx], text.Bold}
	tw := table.NewWriter()
	for _, line := range strings.Split(titleText, "\n") {
		if line != "" {
			tw.AppendRow(table.Row{colors.Sprint(line)})
		}
	}
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}
