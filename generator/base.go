package generator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-sudoku/sudoku"
	"github.com/jedib0t/go-sudoku/writer"
)

type baseGenerator struct {
	attempts               int
	cycles                 int
	debug                  bool
	shouldShowProgress     bool
	showProgressInterval   time.Duration
	showProgressLinesShown int
	rng                    *rand.Rand
}

func (b baseGenerator) Debug() bool {
	return b.debug
}

func (b baseGenerator) RNG() *rand.Rand {
	return b.rng
}

func (b baseGenerator) debugLog(msg string, a ...interface{}) {
	if b.debug {
		_, _ = fmt.Fprintf(os.Stderr, "[DBG] "+strings.TrimSpace(msg)+"\n", a...)
	}
}

func (b *baseGenerator) progressCycleShouldBeShown() bool {
	if b.cycles >= 10000 {
		return b.cycles%500 == 0
	}
	if b.cycles >= 1000 {
		return b.cycles%50 == 0
	}
	if b.cycles >= 100 {
		return b.cycles%5 == 0
	}
	return true
}

func (b *baseGenerator) showProgress(grid *sudoku.Grid, og *sudoku.Grid, forceShow bool) {
	if b.shouldShowProgress && (forceShow || b.progressCycleShouldBeShown()) {
		for b.showProgressLinesShown > 0 && !b.debug {
			fmt.Print(text.CursorUp.Sprint())
			fmt.Print(text.EraseLine.Sprint())
			b.showProgressLinesShown--
		}

		tw := table.NewWriter()
		tw.SetStyle(table.StyleBold)
		switch grid.GetMetadata("type") {
		case "jigsaw":
			if og == nil {
				tw.AppendRow(table.Row{writer.RenderJigSaw(grid)})
			} else {
				tw.AppendRow(table.Row{writer.RenderJigSawDiff(grid, og)})
			}
		default:
			if og == nil {
				tw.AppendRow(table.Row{writer.Render(grid)})
			} else {
				tw.AppendRow(table.Row{writer.RenderDiff(grid, og)})
			}
		}
		tw.AppendFooter(table.Row{fmt.Sprintf("Attempt # %d.%d", b.attempts, b.cycles)})
		tw.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
		})
		tw.Style().Format.Footer = text.FormatDefault
		outStr := tw.Render()
		fmt.Println(outStr)
		b.showProgressLinesShown = strings.Count(outStr, "\n") + 1
		time.Sleep(b.showProgressInterval)
	}
}
