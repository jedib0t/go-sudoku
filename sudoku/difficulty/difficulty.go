package difficulty

import (
	"fmt"
	"strings"
)

var (
	difficultyNameMap = map[Difficulty]string{
		None:   "None",
		Kids:   "Kids",
		Easy:   "Easy",
		Medium: "Medium",
		Hard:   "Hard",
		Insane: "Insane",
	}
)

// Difficulty defines the number of clues available in the Sudoku puzzle. The
// value corresponds to the number of blocks that are pre-filled.
type Difficulty int

// Available Difficulty values.
const (
	None   Difficulty = 81
	Kids   Difficulty = 63
	Easy   Difficulty = 45
	Medium Difficulty = 36
	Hard   Difficulty = 27
	Insane Difficulty = 13
)

// Names returns the names of all supported difficulties.
func Names() []string {
	var rsp []string
	for idx := None; idx > 0; idx-- {
		if name, ok := difficultyNameMap[idx]; ok {
			rsp = append(rsp, name)
		}
	}
	return rsp
}

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
		if strings.ToLower(v) == strLower {
			return k
		}
	}
	return Medium
}
