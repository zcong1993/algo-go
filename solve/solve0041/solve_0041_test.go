package solve0041_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0041"
)

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 2, 0},
			3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0041.FirstMissingPositive(c.nums))
	}
}
