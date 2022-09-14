package main

import (
	"fmt"
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
	colorError     = text.Colors{text.BgRed, text.FgBlack, text.Italic}
	colorShortcuts = text.Colors{text.Italic, text.FgWhite, text.Faint}
	colorTitle     = []text.Color{
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
	statLens           = []int{14, 26, 10}

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

func renderKey(n int, colors [3]text.Colors) string {
	colorBg1 := colors[0]
	colorBg2 := colors[1]
	colorLetter := colors[2]
	key := strings.ToUpper(fmt.Sprint(n))
	if key == "0" {
		key = " "
	}

	return fmt.Sprintf("%s\n%s\n%s",
		colorBg1.Sprint(strings.Repeat("▄", len(key)+2)),
		colorBg2.Sprintf(" %s ", colorLetter.Sprint(key)),
		colorBg1.Sprint(strings.Repeat("▀", len(key)+2)),
	)
}

func renderKeyboard() string {
	var hints map[int]bool
	if *flagHints {
		if p := grid.Possibilities(cursor.X, cursor.Y); p != nil {
			hints = p.AvailableMap()
		}
	}

	tw := table.NewWriter()
	row := table.Row{}
	for n := 1; n <= 9; n++ {
		colors := colorNumbers["key"]
		if grid.CountValue(n) == 9 {
			colors = colorNumbers["key.done"]
		} else if hints[n] {
			colors = colorNumbers["key.hint"]
		}
		row = append(row, renderKey(n, colors))
	}
	tw.AppendRow(row)
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}

func renderShortcuts() string {
	return colorShortcuts.Sprint("help: <H/h> | navigate: ▶ ▲ ▼ ◀  | quit: <Q/q>")
}

func renderStats() string {
	timeGame := time.Now().Sub(timeStart)
	timeGameStr := fmt.Sprintf("%02d:%02d:%02d",
		int(timeGame.Hours()), int(timeGame.Minutes())%60, int(timeGame.Seconds())%60)
	errorMsg := ""
	if errorStr != "" {
		errorMsg = colorError.Sprint(text.AlignCenter.Apply(errorStr, statLens[1]))
	}

	tw := table.NewWriter()
	tw.AppendRow(table.Row{gameMode, errorMsg, timeGameStr})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft, WidthMin: statLens[0], WidthMax: statLens[0], WidthMaxEnforcer: text.Trim},
		{Number: 2, Align: text.AlignCenter, WidthMin: statLens[1], WidthMax: statLens[1], WidthMaxEnforcer: text.Trim},
		{Number: 3, Align: text.AlignRight, WidthMin: statLens[2], WidthMax: statLens[2], WidthMaxEnforcer: text.Trim},
	})
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	tw.Style().Options.DrawBorder = false
	tw.Style().Options.SeparateColumns = false
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
