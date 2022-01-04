package solve0234_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0234"
)

func TestIsPalindrome(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 1, 2, 1})
	assert.False(t, solve0234.IsPalindrome(input))

	input = linkedlist.FromArray([]int{1, 2, 2, 1})
	assert.True(t, solve0234.IsPalindrome(input))
}
