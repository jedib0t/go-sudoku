package samurai

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-sudoku/generator"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/writer"
)

// Generate generates and returns 5 grids that form a Samurai Sudoku. Something
// that should look/render like:
//
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 1 1 1       2 2 2 2 2 2 2 2 2
// 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 2 2 2 2 2 2
// 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 2 2 2 2 2 2
// 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 2 2 2 2 2 2
//             0 0 0 0 0 0 0 0 0
//             0 0 0 0 0 0 0 0 0
//             0 0 0 0 0 0 0 0 0
// 3 3 3 3 3 3 0 0 0 0 0 0 0 0 0 4 4 4 4 4 4
// 3 3 3 3 3 3 0 0 0 0 0 0 0 0 0 4 4 4 4 4 4
// 3 3 3 3 3 3 0 0 0 0 0 0 0 0 0 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
// 3 3 3 3 3 3 3 3 3       4 4 4 4 4 4 4 4 4
func Generate(g generator.Generator) ([]*sudoku.Grid, error) {
	grid0, err := g.Generate(nil)
	if err != nil {
		return nil, fmt.Errorf("error generating middle grid: %v", err)
	}
	grid0.SetMetadata("type", "samurai.0")

	grid1 := sudoku.NewGrid()
	if err := grid1.CopySubGrid(grid0, sudoku.Location{X: 0, Y: 0}, sudoku.Location{X: 6, Y: 6}); err != nil {
		return nil, fmt.Errorf("error cloning top-left subgrid: %v", err)
	}
	grid1, err = g.Generate(grid1)
	if err != nil {
		return nil, fmt.Errorf("error generating top-left grid: %v", err)
	}
	grid1.SetMetadata("type", "samurai.1")

	grid2 := sudoku.NewGrid()
	if err := grid2.CopySubGrid(grid0, sudoku.Location{X: 0, Y: 6}, sudoku.Location{X: 6, Y: 0}); err != nil {
		return nil, fmt.Errorf("error cloning top-right subgrid: %v", err)
	}
	grid2, err = g.Generate(grid2)
	if err != nil {
		return nil, fmt.Errorf("error generating top-right grid: %v", err)
	}
	grid2.SetMetadata("type", "samurai.2")

	grid3 := sudoku.NewGrid()
	if err := grid3.CopySubGrid(grid0, sudoku.Location{X: 6, Y: 0}, sudoku.Location{X: 0, Y: 6}); err != nil {
		return nil, fmt.Errorf("error cloning top-right subgrid: %v", err)
	}
	grid3, err = g.Generate(grid3)
	if err != nil {
		return nil, fmt.Errorf("error generating bottom-left grid: %v", err)
	}
	grid3.SetMetadata("type", "samurai.3")

	grid4 := sudoku.NewGrid()
	if err := grid4.CopySubGrid(grid0, sudoku.Location{X: 6, Y: 6}, sudoku.Location{X: 0, Y: 0}); err != nil {
		return nil, fmt.Errorf("error cloning top-right subgrid: %v", err)
	}
	grid4, err = g.Generate(grid4)
	if err != nil {
		return nil, fmt.Errorf("error generating bottom-right grid: %v", err)
	}
	grid4.SetMetadata("type", "samurai.4")

	// debugging
	if g.Debug() {
		twOriginal := table.NewWriter()
		twOriginal.AppendRow(table.Row{writer.Render(grid1), "", writer.Render(grid2)})
		twOriginal.AppendRow(table.Row{"", writer.Render(grid0), ""})
		twOriginal.AppendRow(table.Row{writer.Render(grid3), "", writer.Render(grid4)})
		twOriginal.SetStyle(table.StyleBold)
		twOriginal.Style().Box.PaddingLeft = ""
		twOriginal.Style().Box.PaddingRight = ""
		twOriginal.Style().Options.SeparateColumns = false
		_, _ = fmt.Fprintln(os.Stderr, twOriginal.Render())
	}

	return []*sudoku.Grid{grid0, grid1, grid2, grid3, grid4}, nil
}
