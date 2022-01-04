package solve0021_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
	"github.com/zcong1993/algo-go/solve/solve0021"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := linkedlist.FromArray([]int{1, 3, 8, 100})
	list2 := linkedlist.FromArray([]int{2, 4, 5, 7, 9})
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 100}

	assert.Equal(t, expected, linkedlist.ToArray(solve0021.MergeTwoLists(list1, list2)))
}
