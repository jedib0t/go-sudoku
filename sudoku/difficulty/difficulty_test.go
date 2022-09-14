package difficulty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	assert.Equal(t, []string{"None", "Kids", "Easy", "Medium", "Hard", "Insane"}, Names())
}

func TestDifficulty_BlocksFilled(t *testing.T) {
	assert.Equal(t, 42, Difficulty(42).BlocksFilled())
	assert.Equal(t, int(None), None.BlocksFilled())
	assert.Equal(t, int(Easy), Easy.BlocksFilled())
	assert.Equal(t, int(Medium), Medium.BlocksFilled())
	assert.Equal(t, int(Hard), Hard.BlocksFilled())
	assert.Equal(t, int(Insane), Insane.BlocksFilled())
}

func TestDifficulty_String(t *testing.T) {
	assert.Equal(t, "None", None.String())
	assert.Equal(t, "Easy", Easy.String())
	assert.Equal(t, "Medium", Medium.String())
	assert.Equal(t, "Hard", Hard.String())
	assert.Equal(t, "Insane", Insane.String())
	assert.Equal(t, "Custom[50]", Difficulty(50).String())
}
