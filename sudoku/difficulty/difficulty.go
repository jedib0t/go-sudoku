package difficulty

import "fmt"

// Difficulty defines the number of clues available in the Sudoku puzzle. The
// value corresponds to the number of blocks that are pre-filled.
type Difficulty int

// Available Difficulty values.
const (
	None   Difficulty = 81
	Easy   Difficulty = 45
	Medium Difficulty = 36
	Hard   Difficulty = 27
	Insane Difficulty = 13
)

// BlocksFilled returns the # of blocks that need to be filled to qualify for
// that difficulty mode.
func (d Difficulty) BlocksFilled() int {
	return int(d)
}

// String returns a human-readable name for the Difficulty.
func (d Difficulty) String() string {
	switch d {
	case None:
		return "None"
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	case Insane:
		return "Insane"
	default:
		return fmt.Sprintf("Custom[%d]", d)
	}
}
