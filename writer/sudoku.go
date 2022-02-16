package writer

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-sudoku/sudoku"
)

// Render renders the given Sudoku Grid in a pretty ASCII table.
func Render(g *sudoku.Grid) string {
	return RenderDiff(g, nil)
}

// RenderDiff renders the given Sudoku Grid in a pretty ASCII table and
// highlights the differences compared to Grid 'og'.
func RenderDiff(g, og *sudoku.Grid) string {
	// no og == no diff, just simple render
	if og == nil {
		og = g
	}

	switch themeSelected {
	case themeNone:
		return renderPlain(g, og)
	default:
		return renderThemed(g, og)
	}
}

func renderPlain(g *sudoku.Grid, og *sudoku.Grid) string {
	sgLocations := [9]sudoku.Location{
		{X: 0, Y: 0}, {X: 0, Y: 3}, {X: 0, Y: 6},
		{X: 3, Y: 0}, {X: 3, Y: 3}, {X: 3, Y: 6},
		{X: 6, Y: 0}, {X: 6, Y: 3}, {X: 6, Y: 6},
	}

	tw := table.NewWriter()
	twRow := table.Row{}
	for _, loc := range sgLocations {
		twSG := table.NewWriter()

		sg := g.SubGrid(loc.X, loc.Y)
		var row table.Row
		for idx, loc := range sg.Locations {
			val := g.Get(loc.X, loc.Y)
			valOG := og.Get(loc.X, loc.Y)

			if val == 0 {
				row = append(row, " ")
			} else if val != valOG {
				row = append(row, colorsDiff.Sprint(val))
			} else {
				row = append(row, fmt.Sprint(val))
			}

			if (idx+1)%3 == 0 {
				twSG.AppendRow(row)
				row = table.Row{}
			}
		}
		//twSG.SetStyle(table.StyleLight)
		twSG.Style().Options = table.OptionsNoBordersAndSeparators
		twRow = append(twRow, twSG.Render())

		if len(twRow) == 3 {
			tw.AppendRow(twRow)
			tw.AppendSeparator()
			twRow = table.Row{}
		}
	}
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.DrawBorder = false
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	return tw.Render()
}

func renderThemed(g *sudoku.Grid, og *sudoku.Grid) string {
	gridIdxRange := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tw := table.NewWriter()
	tw.Style().Box.PaddingRight = ""
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	for rowIdx := range gridIdxRange {
		row := table.Row{}
		for colIdx := range gridIdxRange {
			val := g.Get(rowIdx, colIdx)
			valOG := og.Get(rowIdx, colIdx)

			blockIdx := (rowIdx * 9) + colIdx
			colors := colorsDark
			if gridColorMap[blockIdx] == 1 {
				colors = colorsBright
			}
			if val != valOG {
				colors = append(colors, colorsDiff...)
			}

			if val == 0 {
				row = append(row, colors.Sprintf(" %s ", " "))
			} else {
				row = append(row, colors.Sprintf(" %d ", val))
			}
		}
		tw.AppendRow(row)
	}
	return tw.Render()
}

// Themes returns the list of available themes.
func Themes() []string {
	return themes
}
