package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert.Nil(t, Get("foo"))
	for _, name := range Names() {
		p := Get(name)
		assert.NotNil(t, p)
	}
}

func TestNames(t *testing.T) {
	assert.Len(t, Names(), 6)
}
