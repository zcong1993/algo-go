package solve0054_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0054"
)

func TestSpiralOrder(t *testing.T) {
	inputs := helper.GridFromJSON(`[[1,2,3,4],[5,6,7,8],[9,10,11,12]]`)
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	assert.Equal(t, expected, solve0054.SpiralOrder(inputs))

	inputs = helper.GridFromJSON(`[[1,2,3],[4,5,6],[7,8,9]]`)
	expected = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	assert.Equal(t, expected, solve0054.SpiralOrder(inputs))
}
