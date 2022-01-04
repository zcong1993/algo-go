package solve0015_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0015"
)

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	expected := helper.GridFromJSON(`[[-1, 2, -1], [0, 1, -1]]`)
	assert.Equal(t, expected, solve0015.ThreeSum(input))
}

func TestThreeSum2(t *testing.T) {
	input := []int{1, 1, 1, 2, -3}
	expected := helper.GridFromJSON(`[[1, 2, -3]]`)
	assert.Equal(t, expected, solve0015.ThreeSum(input))
}
