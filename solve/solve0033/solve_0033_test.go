package solve0033_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0033"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		arr      []int
		target   int
		expected int
	}{
		{
			arr:      []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			arr:      []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			arr:      []int{1, 3},
			target:   3,
			expected: 1,
		},
		{
			arr:      []int{3, 1},
			target:   3,
			expected: 0,
		},
		{
			arr:      []int{3},
			target:   3,
			expected: 0,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0033.Search(c.arr, c.target))
	}
}
