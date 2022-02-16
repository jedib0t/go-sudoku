package generator

import (
	"fmt"
	"sort"
	"time"

	"github.com/jedib0t/go-sudoku/sudoku"
)

const (
	backTrackingGeneratorName = "Back-Tracking"
	backTrackingMaxAttempts   = 10000
	backTrackingMaxCycles     = 10000
)

// BackTrackingGenerator generates a Grid with the back-tracking algorithm.
func BackTrackingGenerator(opts ...Option) Generator {
	g := &backTrackingGenerator{}
	opts = append(defaultOptions(), opts...)
	for _, opt := range opts {
		opt(g)
	}
	return g
}

type backTrackingGenerator struct {
	baseGenerator
	blockOrder []int
}

func (b backTrackingGenerator) Name() string {
	return backTrackingGeneratorName
}

func (b backTrackingGenerator) Generate(og *sudoku.Grid) (*sudoku.Grid, error) {
	timeStart := time.Now()
	b.attempts = 0
	b.cycles = 0
	for {
		// abort if this is taking too many attempts or too much time
		if b.attempts > backTrackingMaxAttempts {
			break
		}
		b.attempts++
		b.debugLog("Attempt #%d", b.attempts)

		grid, err := b.generate(og)
		if err == nil {
			b.debugLog("Attempt #%d succeeded", b.attempts)
			grid.SetMetadata("generator", backTrackingGeneratorName)
			grid.SetMetadata("attempts", fmt.Sprint(b.attempts))
			grid.SetMetadata("cycles", fmt.Sprint(b.cycles))
			return grid, nil
		}
		b.debugLog("Attempt #%d failed: %v", err)
	}
	return nil, fmt.Errorf("failed to generate a valid Sudoku after %d attempts in %s", b.attempts, time.Now().Sub(timeStart))
}

func (b *backTrackingGenerator) generate(og *sudoku.Grid) (*sudoku.Grid, error) {
	b.cycles = 0
	grid := cloneOrCreateGrid(og, b.rng)

	// find all the empty blocks, and the possible values for each of the blocks
	emptyBlocks := make([]int, 0)
	possibilitiesMap := make(map[int]*sudoku.Possibilities)
	for idx := 0; idx < 81; idx++ {
		row, col := idx/9, idx%9
		if !grid.IsSet(row, col) {
			emptyBlocks = append(emptyBlocks, idx)
			possibilitiesMap[idx] = grid.Possibilities(row, col)
		}
	}
	// sort empty blocks by the number of possible values in ascending order
	sort.Slice(emptyBlocks, func(i, j int) bool {
		return possibilitiesMap[emptyBlocks[i]].AvailableLen() < possibilitiesMap[emptyBlocks[j]].AvailableLen()
	})
	totalPossibilities := int64(1)
	for _, idx := range emptyBlocks {
		row, col := idx/9, idx%9
		b.debugLog("Possibilities @ #%d [%d, %d]: %v", idx, row, col, possibilitiesMap[idx].Available())
		totalPossibilities *= int64(possibilitiesMap[idx].AvailableLen())
	}
	b.debugLog("Total Possibilities: %d", totalPossibilities)

	// fill all blocks one by one
	direction := 1
	for emptyBlockIdx := 0; emptyBlockIdx < len(emptyBlocks); emptyBlockIdx += direction {
		idx := emptyBlocks[emptyBlockIdx]
		row, col := idx/9, idx%9

		// loop until this block has a valid value set
		for !grid.IsSet(row, col) {
			b.cycles++
			if b.cycles >= backTrackingMaxCycles {
				return nil, fmt.Errorf("aborting due to too many cycles[%d]", b.cycles)
			}

			n := possibilitiesMap[idx].Get()
			if n == 0 { // no more numbers to try
				b.debugLog("[%d] [%d, %d] no more possible values; back-tracking", b.cycles, row, col)
				// reset the current block possibilities
				possibilitiesMap[idx].ResetAvailable()
				// clear out the previous block value
				grid.Reset((emptyBlocks[emptyBlockIdx-1])/9, (emptyBlocks[emptyBlockIdx-1])%9)
				// back-track
				direction = -1
				break
			}

			if grid.Set(row, col, n) {
				b.debugLog("[%d] [%d, %d] set %d successfully", b.cycles, row, col, n)
				b.baseGenerator.showProgress(grid, og, false)
				// success; move ahead
				direction = 1
			}
		}
	}
	b.baseGenerator.showProgress(grid, og, true)

	if _, err := grid.Validate(); err != nil {
		return nil, err
	}
	return grid, nil
}
