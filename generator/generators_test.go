package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerators(t *testing.T) {
	generators := Generators()
	assert.Len(t, generators, 2)
	assert.Equal(t, backTrackingGeneratorName, generators[0])
	assert.Equal(t, bruteForceGeneratorName, generators[1])
}
