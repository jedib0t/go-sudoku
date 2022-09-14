package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/variants/jigsaw"
)

var (
	showProgressLinesShown = 0
	showProgressInterval   = time.Millisecond * 25
)

func renderGeneratedResult(title string, result string) {
	tw := table.NewWriter()
	tw.AppendRow(table.Row{result})
	tw.AppendFooter(table.Row{fmt.Sprintf("%s [%v]", title, getSeed())})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(table.StyleBold)
	tw.Style().Format.Footer = text.FormatDefault
	fmt.Println(tw.Render())
}

func renderProgress(g, og *sudoku.Grid, attempts, cycles int) {
	for showProgressLinesShown > 0 {
		fmt.Print(text.CursorUp.Sprint())
		fmt.Print(text.EraseLine.Sprint())
		showProgressLinesShown--
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleBold)
	switch g.GetMetadata("type") {
	case "jigsaw":
		if og == nil {
			tw.AppendRow(table.Row{jigsaw.Render(g)})
		} else {
			tw.AppendRow(table.Row{jigsaw.RenderDiff(g, og)})
		}
	default:
		if og == nil {
			tw.AppendRow(table.Row{sudoku.Render(g)})
		} else {
			tw.AppendRow(table.Row{sudoku.RenderDiff(g, og)})
		}
	}
	tw.AppendFooter(table.Row{fmt.Sprintf("Attempt # %d.%d", attempts, cycles)})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.Style().Format.Footer = text.FormatDefault
	out := tw.Render()

	fmt.Println(out)
	showProgressLinesShown = strings.Count(out, "\n") + 1
	time.Sleep(showProgressInterval)
}
