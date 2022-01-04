package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/sort"
)

func testSort(t *testing.T, sortFn func(arr []int)) {
	arr := []int{0, 1, 3, 2}
	sortFn(arr)
	expected := []int{0, 1, 2, 3}
	assert.Equal(t, expected, arr)
}

func TestMergeSort(t *testing.T) {
	testSort(t, sort.MergeSort)
}

func TestQuickSort(t *testing.T) {
	testSort(t, sort.QuickSort)
}

func TestQuickSort2(t *testing.T) {
	testSort(t, sort.QuickSort2)
}
