package pattern

// Pattern defines a unique way of making a Sudoku puzzle.
type Pattern [9][9]int

// IsEnabled returns true if the specified block is part of the pattern.
func (p Pattern) IsEnabled(row, col int) bool {
	if row >= 0 && row <= 8 && col >= 0 && col <= 8 {
		return p[row][col] == 1
	}
	return false
}
