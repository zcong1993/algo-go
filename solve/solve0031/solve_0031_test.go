package solve0031_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0031"
)

func TestNextPermutation(t *testing.T) {
	inputs := []int{1, 2, 3}
	expected := []int{1, 3, 2}
	solve0031.NextPermutation(inputs)
	assert.Equal(t, expected, inputs)

	inputs = []int{1, 2, 3, 8, 5, 7, 6, 4}
	expected = []int{1, 2, 3, 8, 6, 4, 5, 7}
	solve0031.NextPermutation(inputs)
	assert.Equal(t, expected, inputs)
}
