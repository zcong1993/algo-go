package solve1236

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 4, Tribonacci(4))
	assert.Equal(t, 149, Tribonacci(10))
	assert.Equal(t, 0, Tribonacci(0))
	assert.Equal(t, 1, Tribonacci(1))
	assert.Equal(t, 1, Tribonacci(2))
}
