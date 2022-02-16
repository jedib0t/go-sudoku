package generator

import (
	"math/rand"
	"sort"
	"time"

	"github.com/jedib0t/go-sudoku/sudoku"
)

var (
	defaultRNG = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Generator defines a method to take Difficulty as input and generate a valid
// and solvable Sudoku Grid.
type Generator interface {
	Debug() bool
	Name() string
	Generate(grid *sudoku.Grid) (*sudoku.Grid, error)
	RNG() *rand.Rand
}

// Generators returns the names of all available Generators.
func Generators() []string {
	rsp := []string{
		BackTrackingGenerator().Name(),
		BruteForceGenerator().Name(),
	}
	sort.Strings(rsp)
	return rsp
}

// Seed seeds the random number Generator. MarkUsed this to reproduce results when
// needed. Default seed is the time at the start of execution.
func Seed(n int64) {
	defaultRNG = rand.New(rand.NewSource(n))
}

// cloneOrCreateGrid returns a fresh Grid which is a copy/clone of the given
// Grid. If none give, gives back a brand-new Grid.
func cloneOrCreateGrid(og *sudoku.Grid, rng *rand.Rand) *sudoku.Grid {
	var grid *sudoku.Grid
	if og != nil {
		grid = og.Clone()
	} else {
		grid = sudoku.NewGrid()
	}
	grid.SetRNG(rng)
	return grid
}
