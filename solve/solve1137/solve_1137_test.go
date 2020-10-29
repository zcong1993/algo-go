package solve1137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 4, Tribonacci(4))
	assert.Equal(t, 149, Tribonacci(10))
}
