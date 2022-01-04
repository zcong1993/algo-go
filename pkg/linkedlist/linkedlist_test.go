package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/linkedlist"
)

func testReverse(t *testing.T, fn func(node *linkedlist.ListNode) *linkedlist.ListNode) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4})
	expected := []int{4, 3, 2, 1}

	assert.Equal(t, expected, linkedlist.ToArray(fn(input)))
}

func testReverseN(t *testing.T, fn func(node *linkedlist.ListNode, n int) *linkedlist.ListNode) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	n := 3
	expected := []int{3, 2, 1, 4, 5}

	assert.Equal(t, expected, linkedlist.ToArray(fn(input, n)))
}

func testReverseNM(t *testing.T, fn func(node *linkedlist.ListNode, n, m int) *linkedlist.ListNode) {
	input := linkedlist.FromArray([]int{1, 2, 3, 4, 5})
	n := 2
	m := 4
	expected := []int{1, 4, 3, 2, 5}

	assert.Equal(t, expected, linkedlist.ToArray(fn(input, n, m)))

	input = linkedlist.FromArray([]int{1, 2})
	n = 1
	m = 2
	expected = []int{2, 1}

	assert.Equal(t, expected, linkedlist.ToArray(fn(input, n, m)))
}

func TestReverse(t *testing.T) {
	testReverse(t, linkedlist.Reverse)
}

func TestReverse2(t *testing.T) {
	testReverse(t, linkedlist.Reverse2)
}

func TestReverseN(t *testing.T) {
	testReverseN(t, linkedlist.ReverseN)
}

func TestReverseN2(t *testing.T) {
	testReverseN(t, linkedlist.ReverseN2)
}

func TestReverseNM(t *testing.T) {
	testReverseNM(t, linkedlist.ReverseNM)
}

func TestReverseNM2(t *testing.T) {
	testReverseNM(t, linkedlist.ReverseNM2)
}

func TestReverseNM3(t *testing.T) {
	testReverseNM(t, linkedlist.ReverseNM3)
}

func TestReverseNM4(t *testing.T) {
	testReverseNM(t, linkedlist.ReverseNM4)
}

func TestCombine(t *testing.T) {
	node1 := linkedlist.FromArray([]int{1, 2, 3, 4})
	node2 := linkedlist.FromArray([]int{5, 6})
	n3 := linkedlist.Combine(node1, node2)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, linkedlist.ToArray(n3))
	assert.Equal(t, []int{5, 6}, linkedlist.ToArray(linkedlist.Combine(nil, node2)))
	assert.Equal(t, []int{5, 6}, linkedlist.ToArray(linkedlist.Combine(node2, nil)))
	assert.Equal(t, []int{}, linkedlist.ToArray(linkedlist.Combine(nil, nil)))
}
