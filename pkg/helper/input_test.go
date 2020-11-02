package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridFromJSON(t *testing.T) {
	expected := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	assert.Equal(t, expected, GridFromJSON("[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]"))
}

func TestArrFromJSON(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	assert.Equal(t, expected, ArrFromJSON("[1,2,3,4]"))
}
