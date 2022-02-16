package sudoku

// Block defines one individual slot in the 9x9 grid.
type Block struct {
	value int
}

// Value returns the currently set value.
func (b *Block) Value() int {
	return b.value
}
