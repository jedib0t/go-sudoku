package sudoku

import (
	"bytes"
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
	"github.com/jedib0t/go-sudoku/sudoku/pattern"
)

var (
	gridIdxRange   = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	subGridIndices = [9][2]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}
)

// Grid is a 3x3 matrix of Grids - basically a 9x9 matrix of integers.
type Grid struct {
	grid     [9][9]Block
	metadata map[string]string
	rng      *rand.Rand
	subGrids SubGrids
}

// NewGrid returns a new Grid object with sane defaults.
func NewGrid(opts ...Option) *Grid {
	rsp := &Grid{}
	for _, opt := range append(defaultOpts, opts...) {
		opt(rsp)
	}
	return rsp
}

// ApplyDifficulty applies the difficulty level on the current grid.
func (g *Grid) ApplyDifficulty(d difficulty.Difficulty) {
	numBlocksToClear := 81 - d.BlocksFilled()
	numClearedCountMap := make(map[int]int, 9)
	for numBlocksCleared := 81 - g.Count(); numBlocksCleared < numBlocksToClear; {
		idx := g.getRNG().Intn(81)
		row, col := idx/9, idx%9
		val := g.Get(row, col)
		if val == 0 { // not set; nothing to clear
			continue
		}
		if g.SubGrid(row, col).Count() == 1 {
			continue // only 1 number left in sub-grid (do not empty)
		}
		if numClearedCountMap[val] == 8 {
			continue // value has been removed 8 times (leave at least one)
		}
		g.Reset(row, col)
		numBlocksCleared++
		numClearedCountMap[val]++
	}
}

// ApplyPattern applies the pattern on the current grid.
func (g *Grid) ApplyPattern(p *pattern.Pattern) {
	if p != nil {
		for row := range gridIdxRange {
			for col := range gridIdxRange {
				if !p.IsEnabled(row, col) {
					g.Reset(row, col)
				}
			}
		}
	}
}

// Checksum returns a unique identifier for the Grid.
func (g Grid) Checksum() string {
	csvString := []byte(g.String())
	sum := sha1.Sum(csvString)
	return fmt.Sprintf("%x", sum)[0:7]
}

// Clone returns a cloned copy of Grid.
func (g Grid) Clone() *Grid {
	g2 := &Grid{rng: g.rng}
	for _, sg := range g.subGrids {
		g2.subGrids = append(g2.subGrids, SubGrid{
			Locations: sg.Locations,
			grid:      g2,
		})
	}
	for row := range gridIdxRange {
		for col := range gridIdxRange {
			g2.Set(row, col, g.Get(row, col))
		}
	}
	for k, v := range g.metadata {
		g2.SetMetadata(k, v)
	}
	return g2
}

// CopySubGrid copies the values from the SubGrid that contains the source
// Location into the SubGrid that contains the target Location.
func (g *Grid) CopySubGrid(srcGrid *Grid, srcLoc, dstLoc Location) error {
	from := srcGrid.SubGrid(srcLoc.X, srcLoc.Y)
	if from == nil {
		return fmt.Errorf("failed to find source sub-grid at [%d, %d]", srcLoc.X, srcLoc.Y)
	}
	to := g.SubGrid(dstLoc.X, dstLoc.Y)
	if to == nil {
		return fmt.Errorf("failed to find destination sub-grid at [%d, %d]", dstLoc.X, dstLoc.Y)
	}
	if len(from.Locations) != len(to.Locations) {
		return fmt.Errorf("from and to sub-grids have differing lengths [%d, %d]",
			len(from.Locations), len(to.Locations))
	}

	for idx, loc := range to.Locations {
		row, col := loc.X, loc.Y
		fromLoc := from.Locations[idx]
		g.Reset(row, col)
		g.Set(row, col, srcGrid.Get(fromLoc.X, fromLoc.Y))
	}
	return nil
}

// Count returns the number of Blocks that are filled and has a value.
func (g Grid) Count() int {
	done := 0
	for row := range gridIdxRange {
		for col := range gridIdxRange {
			if g.grid[row][col].Value() > 0 {
				done++
			}
		}
	}
	return done
}

// CountToDo returns the number of blocks that are not filled and has 0 value.
func (g Grid) CountToDo() int {
	return 81 - g.Count()
}

// Done returns true if all Blocks are filled.
func (g Grid) Done() bool {
	return g.Count() == 81
}

// Empty returns true if all Blocks are empty.
func (g Grid) Empty() bool {
	return g.Count() == 0
}

// Get returns the value stored in the Block at row 'row' and column 'col'.
func (g Grid) Get(row, col int) int {
	return g.grid[row][col].Value()
}

// GetMetadata returns metadata stored against a key.
func (g Grid) GetMetadata(key string) string {
	if g.metadata != nil {
		return g.metadata[key]
	}
	return ""
}

// HasInColumn returns true if the column 'col' contains the number 'n',
func (g Grid) HasInColumn(col, n int) bool {
	for row := range gridIdxRange {
		if g.grid[row][col].Value() == n {
			return true
		}
	}
	return false
}

// HasInRow returns true if the row 'row' contains the number 'n',
func (g Grid) HasInRow(row, n int) bool {
	for col := range gridIdxRange {
		if g.grid[row][col].Value() == n {
			return true
		}
	}
	return false
}

// IsSet returns true if there is a value in the row 'row' and column 'col'.
func (g Grid) IsSet(row, col int) bool {
	return g.Get(row, col) != 0
}

// MarshalArray returns the Grid in a simple 2-dimensional array.
func (g *Grid) MarshalArray() [][]int {
	rsp := make([][]int, 9)
	for row := range gridIdxRange {
		rowNew := make([]int, 9)
		for col := range gridIdxRange {
			rowNew[col] = g.Get(row, col)
		}
		rsp[row] = rowNew
	}
	return rsp
}

// MarshalCSV returns the Grid in a simple CSV string.
func (g *Grid) MarshalCSV() string {
	sb := strings.Builder{}
	sb.Grow(81 + 9)
	for row := range gridIdxRange {
		if row > 0 {
			sb.WriteRune('\n')
		}
		for col := range gridIdxRange {
			if col > 0 {
				sb.WriteRune(',')
			}
			sb.WriteString(fmt.Sprint(g.Get(row, col)))
		}
	}
	return sb.String()
}

// Possibilities returns the Possibilities for the given block at row 'row' and
// column 'col'.
func (g Grid) Possibilities(row, col int) *Possibilities {
	if g.IsSet(row, col) {
		return nil
	}

	possibilities := map[int]bool{
		1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true,
	}
	// remove values set in the row
	for val := range possibilities {
		if g.HasInRow(row, val) {
			delete(possibilities, val)
		}
	}
	// remove values set in the column
	for val := range possibilities {
		if g.HasInColumn(col, val) {
			delete(possibilities, val)
		}
	}
	// remove remaining values set in the sub-grid
	sg := g.SubGrid(row, col)
	for val := range possibilities {
		if sg.Has(val) {
			delete(possibilities, val)
		}
	}
	return NewPossibilitiesFromMap(possibilities, g.rng)
}

// Reset resets the block corresponding to the row 'row' and column 'col'.
func (g *Grid) Reset(row, col int) {
	g.grid[row][col] = Block{value: 0}
}

// Set sets the number 'n' in the row 'row' and column 'col' if it is a valid
// value to set there.
func (g *Grid) Set(row, col, n int) bool {
	if n < 1 || n > 9 {
		return false
	}
	sg := g.SubGrid(row, col)
	if sg.Has(n) || g.HasInRow(row, n) || g.HasInColumn(col, n) {
		return false
	}
	g.grid[row][col] = Block{value: n}
	return true
}

// SetMetadata stores metadata stored against a key.
func (g *Grid) SetMetadata(key, val string) {
	if g.metadata == nil {
		g.metadata = make(map[string]string)
	}
	g.metadata[strings.ToLower(key)] = strings.ToLower(val)
}

// SetRNG sets the RNG to be used by the Grid when needed.
func (g *Grid) SetRNG(rng *rand.Rand) {
	g.rng = rng
}

// String returns the Grid in a human-readable CSV form.
func (g Grid) String() string {
	return g.MarshalCSV()
}

// SubGrid returns the SubGrid corresponding to the given location.
func (g *Grid) SubGrid(row, col int) *SubGrid {
	for _, sg := range g.subGrids {
		if sg.HasLocation(row, col) {
			return &sg
		}
	}
	return nil
}

// SubGrids returns all the registered SubGrids.
func (g *Grid) SubGrids() SubGrids {
	return g.subGrids
}

// UnmarshalArray unmarshals a 2-dimensional array into the Grid.
func (g *Grid) UnmarshalArray(arr [][]int) error {
	if len(arr) != 9 {
		return fmt.Errorf("input does not have 9 rows (%d)", len(arr))
	}
	for rowIdx, row := range arr {
		if len(row) != 9 {
			return fmt.Errorf("row[%d] does not have 9 columns (%d)", rowIdx, len(row))
		}
		for colIdx, val := range row {
			if val < 0 || val > 9 {
				return fmt.Errorf("row[%d] column[%d] contains invalid number '%d'", rowIdx, colIdx, val)
			}
			g.Set(rowIdx, colIdx, val)
		}
	}
	return nil
}

// UnmarshalCSV unmarshals a CSV string into the Grid.
func (g *Grid) UnmarshalCSV(csvString string) error {
	reader := csv.NewReader(bytes.NewReader([]byte(csvString)))
	for rowIdx := range gridIdxRange {
		row, err := reader.Read()
		if err != nil {
			return fmt.Errorf("failed to read row #%d: %v", rowIdx, err)
		}
		if len(row) != 9 {
			return fmt.Errorf("row[%d] does not have 9 columns (%d)", rowIdx, len(row))
		}
		for colIdx, col := range row {
			val, err := strconv.Atoi(col)
			if err != nil {
				return fmt.Errorf("row[%d] column[%d] contains invalid value '%s'", rowIdx, colIdx, col)
			}
			if val < 0 || val > 9 {
				return fmt.Errorf("row[%d] column[%d] contains invalid number '%d'", rowIdx, colIdx, val)
			}
			g.Reset(rowIdx, colIdx)
			g.Set(rowIdx, colIdx, val)
		}
	}
	return nil
}

// Validate validates the Grid and returns if all the values are valid.
func (g *Grid) Validate() (bool, error) {
	// ensure each row has a maximum of 9 unique numbers
	for row := range gridIdxRange {
		numbersFound := make(map[int]int)
		for col := range gridIdxRange {
			val := g.Get(row, col)
			numbersFound[val]++
			if numbersFound[val] > 1 && val != 0 {
				return false, fmt.Errorf("row[%d] has more than one instance of '%d':\n%s", row+1, val, g)
			}
		}
	}
	// ensure each column has a maximum of 9 unique numbers
	for col := range gridIdxRange {
		numbersFound := make(map[int]int)
		for row := range gridIdxRange {
			val := g.Get(row, col)
			numbersFound[val]++
			if numbersFound[val] > 1 && val != 0 {
				return false, fmt.Errorf("column[%d] has more than one instance of '%d':\n%s", col+1, val, g)
			}
		}
	}
	// ensure each sub-grids have a maximum of 9 unique numbers
	for _, sgIdx := range subGridIndices {
		sg := g.SubGrid(sgIdx[0], sgIdx[1])
		if _, err := sg.Validate(); err != nil {
			return false, fmt.Errorf("sub-grid at [%d, %d] is invalid: %v:\n%s",
				sgIdx[0]+1, sgIdx[1]+1, err, g)
		}
	}
	return true, nil
}

func (g *Grid) getRNG() *rand.Rand {
	if g.rng == nil {
		g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return g.rng
}
