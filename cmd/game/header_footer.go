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
	colorShortcuts = text.Colors{text.Italic, text.FgWhite}
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
	tw := table.NewWriter()
	row := table.Row{}
	for idx := 1; idx <= 9; idx++ {
		row = append(row, renderKey(idx, colorNumbers["key"]))
	}
	tw.AppendRow(row)
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}

func renderShortcuts() string {
	return colorShortcuts.Sprint("insert: 1-9 | navigate: ▶ ▲ ▼ ◀  | quit: <Q/q>")
}

func renderStats() string {
	difficulty := text.FormatTitle.Apply(diff.String())
	timeGame := time.Now().Sub(timeStart)
	timeGameStr := fmt.Sprintf("%02d:%02d:%02d",
		int(timeGame.Hours()), int(timeGame.Minutes()), int(timeGame.Seconds()))
	errorMsg := ""
	if errorStr != "" {
		errorMsg = colorError.Sprint(text.AlignCenter.Apply(errorStr, 30))
	}

	tw := table.NewWriter()
	tw.AppendRow(table.Row{difficulty, errorMsg, timeGameStr})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft, WidthMin: 9, WidthMax: 9, WidthMaxEnforcer: text.Trim},
		{Number: 2, Align: text.AlignCenter, WidthMin: 30, WidthMax: 30, WidthMaxEnforcer: text.Trim},
		{Number: 3, Align: text.AlignRight, WidthMin: 10, WidthMax: 10, WidthMaxEnforcer: text.Trim},
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
