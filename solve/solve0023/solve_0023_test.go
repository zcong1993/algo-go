package solve0023_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0023"
)

func TestMergeKLists(t *testing.T) {
	list1 := linkedlist.FromArray([]int{1, 4, 5})
	list2 := linkedlist.FromArray([]int{1, 3, 4})
	list3 := linkedlist.FromArray([]int{2, 6})
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}

	res := solve0023.MergeKLists([]*linkedlist.ListNode{list1, list2, list3})

	assert.Equal(t, expected, linkedlist.ToArray(res))
}
