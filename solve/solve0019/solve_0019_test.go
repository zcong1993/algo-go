package solve0019_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0019"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		input    []int
		n        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
	}

	for _, c := range cases {
		res := solve0019.RemoveNthFromEnd(linkedlist.FromArray(c.input), c.n)
		assert.Equal(t, c.expected, linkedlist.ToArray(res))
	}
}
