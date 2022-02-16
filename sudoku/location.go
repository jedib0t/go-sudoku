package sudoku

// Location points to a specific location on the Grid (x, y).
type Location struct {
	X int
	Y int
}

// IsValid returns true is the Location points to a location within a 9x9 Grid.
func (l Location) IsValid() bool {
	return l.X >= 0 && l.X <= 8 && l.Y >= 0 && l.Y <= 8
}
