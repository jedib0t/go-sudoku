package utils

import (
	"math/rand"
)

// RandomNumberSequencer defines interfaces for a Sequencer that provides Random
// Numbers using an RNG.
type RandomNumberSequencer interface {
	Get() int
	IsUsed(n int) bool
	MarkUsed(n int)
}

type randomNumberSequencer struct {
	max    int
	rng    *rand.Rand
	values map[int]bool
}

// NewRandomNumberSequencer returns a new RandomNumberSequencer that uses the
// provided RNG and returns a maximum of 'max' numbers after which it starts
// returning 0.
func NewRandomNumberSequencer(rng *rand.Rand, max int) RandomNumberSequencer {
	return &randomNumberSequencer{
		max:    max,
		rng:    rng,
		values: make(map[int]bool, max),
	}
}

// Get returns the next random number in the sequence. If all possible values
// have been exhausted, it starts returning 0.
func (rns randomNumberSequencer) Get() int {
	triedValues := make(map[int]bool)
	for len(triedValues) < rns.max {
		n := rns.rng.Intn(rns.max) + 1
		if !rns.IsUsed(n) {
			return n
		}
		triedValues[n] = true
	}
	return 0
}

// IsUsed returns true if the given number 'n' was called with MarkUsed.
func (rns randomNumberSequencer) IsUsed(n int) bool {
	return rns.values[n]
}

// MarkUsed marks the number 'n' as used, and will result in Get not returning
// this number anymore.
func (rns randomNumberSequencer) MarkUsed(n int) {
	rns.values[n] = true
}
