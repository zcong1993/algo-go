package solve0142_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0142"
)

func TestDetectCycle(t *testing.T) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4})
	var node2 *linkedlist.ListNode
	var tail *linkedlist.ListNode
	cur := input
	i := 1
	for cur != nil {
		if i == 2 {
			node2 = cur
		}
		tail = cur
		cur = cur.Next
		i++
	}
	tail.Next = node2

	res := solve0142.DetectCycle(input)
	assert.Equal(t, 2, res.Val)
}
