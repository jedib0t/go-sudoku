package utils

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumberSequencer(t *testing.T) {
	pseudoRNG := rand.New(rand.NewSource(1))
	rns := NewRandomNumberSequencer(pseudoRNG, 5)
	expectedValues := []int{2, 3, 5, 4, 1, 0}

	for idx, expectedVal := range expectedValues {
		val := rns.Get()
		assert.Equal(t, expectedVal, val, idx)
		rns.MarkUsed(val)
	}
}
