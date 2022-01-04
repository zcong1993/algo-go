package solve0092_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0092"
)

func TestReverseBetween(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	expected := []int{1, 4, 3, 2, 5}

	assert.Equal(t, expected, linkedlist.ToArray(solve0092.ReverseBetween(input, 2, 4)))
}

func TestReverseBetween2(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	expected := []int{1, 4, 3, 2, 5}

	assert.Equal(t, expected, linkedlist.ToArray(solve0092.ReverseBetween2(input, 2, 4)))
}

func TestReverseBetween3(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	expected := []int{1, 4, 3, 2, 5}

	assert.Equal(t, expected, linkedlist.ToArray(solve0092.ReverseBetween3(input, 2, 4)))
}
