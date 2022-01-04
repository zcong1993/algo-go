package solve0056_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0056"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			helper.GridFromJSON(`[[1,3],[2,6],[8,10],[15,18]]`),
			helper.GridFromJSON(`[[1,6],[8,10],[15,18]]`),
		},
		{
			helper.GridFromJSON(`[[1,4],[4,5]]`),
			helper.GridFromJSON(`[[1,5]]`),
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0056.Merge(c.input))
	}
}
