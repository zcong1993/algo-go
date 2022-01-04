package solve0103_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0103"
)

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	solve0103.Reverse(input)
	assert.Equal(t, expected, input)
}

func TestZigzagLevelOrder(t *testing.T) {
	input := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	expected := helper.GridFromJSON(`[
  [3],
  [20,9],
  [15,7]
]`)
	assert.Equal(t, expected, solve0103.ZigzagLevelOrder(input))
}
