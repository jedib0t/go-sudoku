package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPattern_IsEnabled(t *testing.T) {
	p := Get("triangle")
	assert.NotNil(t, p)
	if p != nil {
		assert.False(t, p.IsEnabled(0, 0))
		assert.False(t, p.IsEnabled(1, 1))
		assert.False(t, p.IsEnabled(2, 2))
		assert.False(t, p.IsEnabled(3, 3))
		assert.True(t, p.IsEnabled(4, 4))
		assert.True(t, p.IsEnabled(5, 5))
		assert.True(t, p.IsEnabled(6, 6))
		assert.True(t, p.IsEnabled(7, 7))
		assert.True(t, p.IsEnabled(8, 8))
	}
}
