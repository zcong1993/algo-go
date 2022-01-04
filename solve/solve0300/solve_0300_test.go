package solve0300_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0300"
)

func TestLengthOfLIS(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{1, 1, 1, 1, 1},
			1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0300.LengthOfLIS(c.nums))
	}
}
