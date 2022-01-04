package solve0792_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0792"
)

func TestSearch(t *testing.T) {
	input := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	assert.Equal(t, expected, solve0792.Search(input, target))

	input = []int{-1, 0, 3, 5, 9, 12}
	target = 2
	expected = -1
	assert.Equal(t, expected, solve0792.Search(input, target))
}
