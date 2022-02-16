package writer

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
)

var (
	jigSawColors = [9]text.Colors{
		{text.BgHiBlack, text.FgWhite},
		{text.BgHiBlue, text.FgBlack},
		{text.BgHiCyan, text.FgBlack},
		{text.BgHiGreen, text.FgBlack},
		{text.BgHiMagenta, text.FgBlack},
		{text.BgHiRed, text.FgBlack},
		{text.BgHiWhite, text.FgBlack},
		{text.BgHiYellow, text.FgBlack},
		{text.BgYellow, text.FgBlack},
	}
	jigSawColorDiff = text.Colors{text.Italic}
)

// RenderJigSaw renders a JigSaw Sudoku in colored mode.
func RenderJigSaw(g *sudoku.Grid) string {
	subGrids := g.SubGrids()

	tw := table.NewWriter()
	for row := 0; row < 9; row++ {
		twRow := table.Row{}
		for col := 0; col < 9; col++ {
			colors := getSubGridColor(subGrids, row, col)
			val := fmt.Sprint(g.Get(row, col))
			if val == "0" {
				val = " "
			}
			twRow = append(twRow, colors.Sprintf(" %s ", val))
		}
		tw.AppendRow(twRow)
	}
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}

// RenderJigSawDiff renders a JigSaw Sudoku in colored mode and highlights the
// differences compared to Grid 'og'.
func RenderJigSawDiff(g, og *sudoku.Grid) string {
	subGrids := g.SubGrids()

	tw := table.NewWriter()
	for row := 0; row < 9; row++ {
		twRow := table.Row{}
		for col := 0; col < 9; col++ {
			colors := getSubGridColor(subGrids, row, col)
			val := g.Get(row, col)
			valOG := og.Get(row, col)
			if val != valOG {
				colors = append(colors, jigSawColorDiff...)
			}
			if val == 0 {
				twRow = append(twRow, colors.Sprint("   "))
			} else {
				twRow = append(twRow, colors.Sprintf(" %d ", val))
			}
		}
		tw.AppendRow(twRow)
	}
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}

func getSubGridColor(subGrids []sudoku.SubGrid, row, col int) text.Colors {
	for idx, sg := range subGrids {
		if idx < len(jigSawColors) && sg.HasLocation(row, col) {
			return jigSawColors[idx]
		}
	}
	return text.Colors{text.FgHiRed}
}
