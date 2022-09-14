package difficulty

import (
	"fmt"
	"strings"
)

var (
	difficultyNameMap = map[Difficulty]string{
		None:   "none",
		Easy:   "easy",
		Medium: "medium",
		Hard:   "hard",
		Insane: "insane",
	}
)

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
	if name, ok := difficultyNameMap[d]; ok && name != "" {
		return name
	}
	return fmt.Sprintf("Custom[%d]", d)
}

// From interprets the given difficulty in string form to the actual enum value.
func From(str string) Difficulty {
	strLower := strings.ToLower(str)
	for k, v := range difficultyNameMap {
		if v == strLower {
			return k
		}
	}
	return Medium
}
