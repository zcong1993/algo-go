package solve0121_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0121"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices   []int
		expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0121.MaxProfit(c.prices))
	}
}
