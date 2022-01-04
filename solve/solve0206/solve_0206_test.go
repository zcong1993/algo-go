package solve0206_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0206"
)

func TestReverseList(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4})
	expected := []int{4, 3, 2, 1}
	assert.Equal(t, expected, linkedlist.ToArray(solve0206.ReverseList(input)))
}

func TestReverseList2(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4})
	expected := []int{4, 3, 2, 1}
	assert.Equal(t, expected, linkedlist.ToArray(solve0206.ReverseList2(input)))
}
