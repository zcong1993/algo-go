package solve0057

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {3, 10}, {12, 16}}, Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	assert.Equal(t, [][]int{{1, 5}}, Insert([][]int{{1, 5}}, []int{2, 3}))
	assert.Equal(t, [][]int{{1, 7}}, Insert([][]int{{1, 5}}, []int{2, 7}))
}
