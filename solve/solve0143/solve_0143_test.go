package solve0143_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0143"
)

func TestReorderList(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	expected := []int{1, 5, 2, 4, 3}
	solve0143.ReorderList(input)
	assert.Equal(t, expected, linkedlist.ToArray(input))

	input = linkedlist.FromArray([]int{1, 2, 3, 4})
	expected = []int{1, 4, 2, 3}
	solve0143.ReorderList(input)
	assert.Equal(t, expected, linkedlist.ToArray(input))
}
