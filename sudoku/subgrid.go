package sudoku

import (
	"fmt"
	"strings"
)

// SubGrid is a 3x3 matrix of integers.
type SubGrid struct {
	Locations []Location
	grid      *Grid
}

// Count returns the number of Blocks that are filled and has a value.
func (s SubGrid) Count() int {
	if s.grid == nil {
		return 0
	}
	done := 0
	for _, loc := range s.Locations {
		if s.grid.Get(loc.X, loc.Y) > 0 {
			done++
		}
	}
	return done
}

// Done returns true if all Blocks are filled.
func (s SubGrid) Done() bool {
	return s.Count() == 9
}

// Empty returns true if all Blocks are empty.
func (s SubGrid) Empty() bool {
	return s.Count() == 0
}

// Has returns true if the SubGrid has the given number in some location.
func (s SubGrid) Has(n int) bool {
	if s.grid == nil {
		return false
	}
	for _, loc := range s.Locations {
		if s.grid.Get(loc.X, loc.Y) == n {
			return true
		}
	}
	return false
}

// HasLocation returns true if the SubGrid has the given location.
func (s SubGrid) HasLocation(x, y int) bool {
	for _, loc := range s.Locations {
		if x == loc.X && y == loc.Y {
			return true
		}
	}
	return false
}

// String returns the SubGrid in a human-readable CSV form.
func (s SubGrid) String() string {
	if s.grid == nil {
		return ""
	}

	sb := &strings.Builder{}
	sb.WriteRune('[')
	for idx, loc := range s.Locations {
		if idx > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(fmt.Sprint(s.grid.Get(loc.X, loc.Y)))
	}
	sb.WriteRune(']')
	return sb.String()
}

// Validate validates the SubGrid and returns if all the values are valid.
func (s SubGrid) Validate() (bool, error) {
	numbersFound := make(map[int]int)
	for _, loc := range s.Locations {
		val := s.grid.Get(loc.X, loc.Y)
		numbersFound[val]++
		if numbersFound[val] > 1 && val != 0 {
			return false, fmt.Errorf("more than one instance of '%d'", val)
		}
	}
	return true, nil
}

// SubGrids represent a list of WithSubGrids in a 9x9 Grid
type SubGrids []SubGrid

// String returns the WithSubGrids in a human-readable format.
func (sgs SubGrids) String() string {
	var grid [9][9]int
	for sgIdx, sg := range sgs {
		for _, loc := range sg.Locations {
			grid[loc.X][loc.Y] = sgIdx + 1
		}
	}
	sb := strings.Builder{}
	for row := range grid {
		if row > 0 {
			sb.WriteRune('\n')
		}
		for col := range grid[row] {
			if col > 0 {
				sb.WriteRune(' ')
			}
			sb.WriteString(fmt.Sprint(grid[row][col]))
		}
	}
	return sb.String()
}
