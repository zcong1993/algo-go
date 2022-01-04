package solve0215_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0215"
)

func TestFindKthLargest(t *testing.T) {
	input := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5
	assert.Equal(t, expected, solve0215.FindKthLargest(input, k))
}

func TestFindKthLargest2(t *testing.T) {
	input := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5
	assert.Equal(t, expected, solve0215.FindKthLargest2(input, k))
}
