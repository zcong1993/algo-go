package solve0160_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0160"
)

func TestGetIntersectionNode(t *testing.T) {
	node3 := linkedlist.FromArray([]int{9, 10})
	node1 := linkedlist.Combine(linkedlist.FromArray([]int{1, 2, 3}), node3)
	node2 := linkedlist.Combine(linkedlist.FromArray([]int{4, 5}), node3)

	node := solve0160.GetIntersectionNode(node1, node2)
	assert.Equal(t, 9, node.Val)
}

func TestGetIntersectionNode2(t *testing.T) {
	node3 := linkedlist.FromArray([]int{9, 10})
	node1 := linkedlist.Combine(linkedlist.FromArray([]int{1, 2, 3}), node3)
	node2 := linkedlist.Combine(linkedlist.FromArray([]int{4, 5}), node3)

	node := solve0160.GetIntersectionNode2(node1, node2)
	assert.Equal(t, 9, node.Val)
}
