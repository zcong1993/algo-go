package solve0025_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0025"
)

func TestReverseKGroup(t *testing.T) {
	cases := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			k:        2,
			expected: []int{2, 1, 4, 3},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:        3,
			expected: []int{3, 2, 1, 6, 5, 4, 7, 8},
		},
	}

	for _, c := range cases {
		res := solve0025.ReverseKGroup(linkedlist.FromArray(c.input), c.k)
		assert.Equal(t, c.expected, linkedlist.ToArray(res))
	}
}
