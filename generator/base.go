package generator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/jedib0t/go-sudoku/sudoku"
)

type baseGenerator struct {
	attempts         int
	cycles           int
	debug            bool
	renderProgressFn func(g, og *sudoku.Grid, attempts, cycles int)
	rng              *rand.Rand
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
	if b.renderProgressFn != nil && (forceShow || b.progressCycleShouldBeShown()) {
		b.renderProgressFn(grid, og, b.attempts, b.cycles)
	}
}
