package samurai

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
)

var (
	colorsBright = text.Colors{}
	colorsDark   = text.Colors{}
)

// Render renders a Samurai Sudoku which is a combination of 5 Sudokus.
func Render(grids []*sudoku.Grid) (string, error) {
	if len(grids) != 5 {
		return " ", fmt.Errorf("found %d interconnected grids instead of 5", len(grids))
	}

	tw := table.NewWriter()
	tw.AppendRow(table.Row{
		renderSubGrid(grids[1], sudoku.Location{X: 0, Y: 0}, colorsDark),
		renderSubGrid(grids[1], sudoku.Location{X: 0, Y: 3}, colorsBright),
		renderSubGrid(grids[1], sudoku.Location{X: 0, Y: 6}, colorsDark),
		" ",
		renderSubGrid(grids[2], sudoku.Location{X: 0, Y: 0}, colorsDark),
		renderSubGrid(grids[2], sudoku.Location{X: 0, Y: 3}, colorsBright),
		renderSubGrid(grids[2], sudoku.Location{X: 0, Y: 6}, colorsDark),
	})
	tw.AppendRow(table.Row{
		renderSubGrid(grids[1], sudoku.Location{X: 3, Y: 0}, colorsBright),
		renderSubGrid(grids[1], sudoku.Location{X: 3, Y: 3}, colorsDark),
		renderSubGrid(grids[1], sudoku.Location{X: 3, Y: 6}, colorsBright),
		" ",
		renderSubGrid(grids[2], sudoku.Location{X: 3, Y: 0}, colorsBright),
		renderSubGrid(grids[2], sudoku.Location{X: 3, Y: 3}, colorsDark),
		renderSubGrid(grids[2], sudoku.Location{X: 3, Y: 6}, colorsBright),
	})
	tw.AppendRow(table.Row{
		renderSubGrid(grids[1], sudoku.Location{X: 6, Y: 0}, colorsDark),
		renderSubGrid(grids[1], sudoku.Location{X: 6, Y: 3}, colorsBright),
		renderSubGrid(grids[0], sudoku.Location{X: 0, Y: 0}, colorsDark),
		renderSubGrid(grids[0], sudoku.Location{X: 0, Y: 3}, colorsBright),
		renderSubGrid(grids[0], sudoku.Location{X: 0, Y: 6}, colorsDark),
		renderSubGrid(grids[2], sudoku.Location{X: 6, Y: 3}, colorsBright),
		renderSubGrid(grids[2], sudoku.Location{X: 6, Y: 6}, colorsDark),
	})
	tw.AppendRow(table.Row{
		" ",
		" ",
		renderSubGrid(grids[0], sudoku.Location{X: 3, Y: 0}, colorsBright),
		renderSubGrid(grids[0], sudoku.Location{X: 3, Y: 3}, colorsDark),
		renderSubGrid(grids[0], sudoku.Location{X: 3, Y: 6}, colorsBright),
		" ",
		" ",
	}, table.RowConfig{AutoMerge: true})
	tw.AppendRow(table.Row{
		renderSubGrid(grids[3], sudoku.Location{X: 0, Y: 0}, colorsDark),
		renderSubGrid(grids[3], sudoku.Location{X: 0, Y: 3}, colorsBright),
		renderSubGrid(grids[0], sudoku.Location{X: 6, Y: 0}, colorsDark),
		renderSubGrid(grids[0], sudoku.Location{X: 6, Y: 3}, colorsBright),
		renderSubGrid(grids[0], sudoku.Location{X: 6, Y: 6}, colorsDark),
		renderSubGrid(grids[4], sudoku.Location{X: 0, Y: 3}, colorsBright),
		renderSubGrid(grids[4], sudoku.Location{X: 0, Y: 6}, colorsDark),
	})
	tw.AppendRow(table.Row{
		renderSubGrid(grids[3], sudoku.Location{X: 3, Y: 0}, colorsBright),
		renderSubGrid(grids[3], sudoku.Location{X: 3, Y: 3}, colorsDark),
		renderSubGrid(grids[3], sudoku.Location{X: 3, Y: 6}, colorsBright),
		" ",
		renderSubGrid(grids[4], sudoku.Location{X: 3, Y: 0}, colorsBright),
		renderSubGrid(grids[4], sudoku.Location{X: 3, Y: 3}, colorsDark),
		renderSubGrid(grids[4], sudoku.Location{X: 3, Y: 6}, colorsBright),
	})
	tw.AppendRow(table.Row{
		renderSubGrid(grids[3], sudoku.Location{X: 6, Y: 0}, colorsDark),
		renderSubGrid(grids[3], sudoku.Location{X: 6, Y: 3}, colorsBright),
		renderSubGrid(grids[3], sudoku.Location{X: 6, Y: 6}, colorsDark),
		" ",
		renderSubGrid(grids[4], sudoku.Location{X: 6, Y: 0}, colorsDark),
		renderSubGrid(grids[4], sudoku.Location{X: 6, Y: 3}, colorsBright),
		renderSubGrid(grids[4], sudoku.Location{X: 6, Y: 6}, colorsDark),
	})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 4, AutoMerge: true},
	})
	tw.SetStyle(table.StyleLight)
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	tw.Style().Options.DrawBorder = false
	tw.Style().Options.SeparateColumns = true
	tw.Style().Options.SeparateRows = true
	return tw.Render(), nil
}

func renderSubGrid(grid *sudoku.Grid, sgLoc sudoku.Location, colors text.Colors) string {
	tw := table.NewWriter()
	sg := grid.SubGrid(sgLoc.X, sgLoc.Y)
	var row table.Row
	for idx, loc := range sg.Locations {
		val := grid.Get(loc.X, loc.Y)
		if val == 0 {
			row = append(row, colors.Sprint("   "))
		} else {
			row = append(row, colors.Sprintf(" %d ", val))
		}
		if (idx+1)%3 == 0 {
			tw.AppendRow(row)
			row = table.Row{}
		}
	}
	tw.SetStyle(table.StyleLight)
	tw.Style().Box.PaddingLeft = ""
	tw.Style().Box.PaddingRight = ""
	tw.Style().Options.DrawBorder = false
	tw.Style().Options.SeparateColumns = false
	return tw.Render()
}
