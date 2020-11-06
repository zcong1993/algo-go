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

func TestStringBoardFromJSON(t *testing.T) {
	expected := [][]string{{"a", "b"}, {"c", "d"}}
	assert.Equal(t, expected, StringBoardFromJSON(`[["a","b"],["c", "d"]]`))
}

func TestCharBoardFromJSON(t *testing.T) {
	expected := [][]byte{{'a', 'b'}, {'c', 'd'}}
	assert.Equal(t, expected, CharBoardFromJSON(`[["a","b"],["c", "d"]]`))
}
