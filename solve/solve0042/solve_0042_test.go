package solve0042_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0042"
)

func TestTrap(t *testing.T) {
	inputs := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	assert.Equal(t, expected, solve0042.Trap(inputs))

	inputs = []int{4, 2, 0, 3, 2, 5}
	expected = 9
	assert.Equal(t, expected, solve0042.Trap(inputs))
}
