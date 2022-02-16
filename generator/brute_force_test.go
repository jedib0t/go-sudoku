package generator

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBruteForceGenerator_Generate(t *testing.T) {
	rng := rand.New(rand.NewSource(13))
	grid, err := BruteForceGenerator(WithRNG(rng)).Generate(nil)
	assert.NotNil(t, grid)
	assert.Nil(t, err)

	expectedGridCSV := `2,3,6,4,7,1,5,9,8
5,8,4,9,2,6,7,3,1
9,7,1,8,3,5,6,4,2
4,5,2,6,1,7,3,8,9
6,1,8,3,9,2,4,7,5
7,9,3,5,8,4,1,2,6
8,6,7,1,4,9,2,5,3
1,4,9,2,5,3,8,6,7
3,2,5,7,6,8,9,1,4`
	assert.Equal(t, expectedGridCSV, grid.MarshalCSV())
}

func TestBruteForceGenerator_Name(t *testing.T) {
	assert.Equal(t, bruteForceGeneratorName, BruteForceGenerator().Name())
}
