package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPossibilitiesFromMap(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false}, testRNG)
	assert.NotNil(t, p)
	assert.Len(t, p.available, 0)
	assert.Equal(t, p.rng, testRNG)
	assert.Len(t, p.possibilities, 0)
}

func TestPossibilities_Available(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false, 2: true, 3: true}, testRNG)
	assert.NotNil(t, p)
	assert.Equal(t, []int{2, 3}, p.Available())
}

func TestPossibilities_AvailableMap(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false, 2: true, 3: true}, testRNG)
	assert.NotNil(t, p)
	assert.Equal(t, map[int]bool{2: true, 3: true}, p.AvailableMap())
}

func TestPossibilities_AvailableLen(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false, 2: true, 3: true}, testRNG)
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.AvailableLen())
}

func TestPossibilities_Get(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false, 2: true, 3: true}, testRNG)
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.AvailableLen())
	assert.Equal(t, 2, p.Get())
	assert.Equal(t, 1, p.AvailableLen())
	assert.Equal(t, 3, p.Get())
	assert.Equal(t, 0, p.AvailableLen())
	assert.Equal(t, 0, p.Get())
	assert.Equal(t, 0, p.AvailableLen())
}

func TestPossibilities_ResetAvailable(t *testing.T) {
	p := NewPossibilitiesFromMap(map[int]bool{1: false, 2: true, 3: true}, testRNG)
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.Get())
	assert.Equal(t, 3, p.Get())
	assert.Equal(t, 0, p.AvailableLen())
	assert.Empty(t, p.available)
	p.ResetAvailable()
	assert.NotEmpty(t, p.available)
	assert.Equal(t, 2, p.Get())
	assert.Equal(t, 3, p.Get())
	assert.Equal(t, 0, p.AvailableLen())
	assert.Empty(t, p.available)
}
