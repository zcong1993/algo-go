package solve0908_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0908"
)

func TestMiddleNode(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	expected := []int{3, 4, 5}
	assert.Equal(t, expected, linkedlist.ToArray(solve0908.MiddleNode(input)))
}
