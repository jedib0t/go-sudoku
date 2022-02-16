package sudoku

import (
	"math/rand"
	"sort"
)

// Possibilities keep track of possible values for a single Block.
type Possibilities struct {
	available     []int
	possibilities map[int]bool
	rng           *rand.Rand
}

// NewPossibilitiesFromMap returns a Possibilities object based on the given map. Key
// being the number (1-9) and the value being whether it is possible or not.
func NewPossibilitiesFromMap(data map[int]bool, rng *rand.Rand) *Possibilities {
	for k, v := range data {
		if !v {
			delete(data, k)
		}
	}
	rsp := &Possibilities{
		possibilities: data,
		rng:           rng,
	}
	rsp.ResetAvailable()
	return rsp
}

// Available returns the available values.
func (p Possibilities) Available() []int {
	return p.available
}

// AvailableLen returns the number of available values.
func (p Possibilities) AvailableLen() int {
	return len(p.available)
}

// Get returns the first possible number and takes it out of available numbers
// list.
func (p *Possibilities) Get() int {
	if len(p.available) == 0 {
		return 0
	}
	idx := p.rng.Intn(len(p.available))
	rsp := p.available[idx]
	p.available = append(p.available[:idx], p.available[idx+1:]...)
	return rsp
}

// ResetAvailable resets the list of available numbers.
func (p *Possibilities) ResetAvailable() {
	p.available = nil
	for k, v := range p.possibilities {
		if v {
			p.available = append(p.available, k)
		}
	}
	sort.Ints(p.available)
}
