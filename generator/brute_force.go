package generator

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-sudoku/sudoku"
)

const (
	bruteForceGeneratorName = "Brute-Force"
	bruteForceMaxAttempts   = 100000
)

// BruteForceGenerator generates a Grid with brute-force. There is no
// intelligent algorithm to it and this has a chance to fail and hence the
// obscene amount of retries that are attempted.
func BruteForceGenerator(opts ...Option) Generator {
	g := &bruteForceGenerator{}
	opts = append(defaultOptions(), opts...)
	for _, opt := range opts {
		opt(g)
	}
	return g
}

type bruteForceGenerator struct {
	baseGenerator
	subGridOrder [][]int
}

func (b bruteForceGenerator) Name() string {
	return bruteForceGeneratorName
}

func (b *bruteForceGenerator) Generate(og *sudoku.Grid) (*sudoku.Grid, error) {
	timeStart := time.Now()
	b.attempts = 0
	b.cycles = 0
	for {
		// abort if this is taking too many attempts or too much time
		if b.attempts > bruteForceMaxAttempts {
			break
		}
		b.attempts++
		b.debugLog("Attempt #%d", b.attempts)

		// generate numbers for each sub-grids one-by-one
		grid, err := b.generateByNumberAndSubGrids(cloneOrCreateGrid(og, b.rng))
		if err == nil {
			b.debugLog("Attempt #%d succeeded", b.attempts)
			grid.SetMetadata("generator", bruteForceGeneratorName)
			grid.SetMetadata("attempts", fmt.Sprint(b.attempts))
			grid.SetMetadata("cycles", fmt.Sprint(b.cycles))
			return grid, nil
		}
		b.debugLog("Attempt #%d failed: %v", err)
	}
	return nil, fmt.Errorf("failed to generate a valid Sudoku after %d attempts in %s", b.attempts, time.Now().Sub(timeStart))
}

func (b *bruteForceGenerator) generateByNumberAndSubGrids(grid *sudoku.Grid) (*sudoku.Grid, error) {
	b.cycles = 0

	rns := NewRandomNumberSequencer(defaultRNG, 9)
	for n := rns.Get(); n > 0; n = rns.Get() {
		for _, xy := range b.subGridOrder {
			err := b.generateForSubGrid(grid, xy[0], xy[1], n)
			if err != nil {
				return nil, err
			}
			b.baseGenerator.showProgress(grid, nil, false)
		}
		rns.MarkUsed(n)
	}

	if _, err := grid.Validate(); err != nil {
		return nil, err
	}
	b.baseGenerator.showProgress(grid, nil, true)
	return grid, nil
}

func (b *bruteForceGenerator) generateForSubGrid(grid *sudoku.Grid, row, col, n int) error {
	defer func() {
		b.cycles++
	}()
	rnsRow := NewRandomNumberSequencer(defaultRNG, 3)
	for sgRow := rnsRow.Get(); sgRow > 0; sgRow = rnsRow.Get() {
		rnsCol := NewRandomNumberSequencer(defaultRNG, 3)
		for sgCol := rnsCol.Get(); sgCol > 0; sgCol = rnsCol.Get() {
			fRow, fCol := row+sgRow-1, col+sgCol-1
			if !grid.IsSet(fRow, fCol) {
				if grid.Set(fRow, fCol, n) {
					b.debugLog("[%d] [%d, %d] set %d successfully", b.cycles, fRow, fCol, n)
					return nil
				}
			}
			rnsCol.MarkUsed(sgCol)
		}
		rnsRow.MarkUsed(sgRow)
	}
	return fmt.Errorf("failed to find a block in sub-grid [%d, %d] for %d", row, col, n)
}
