package solve0200_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0200"
)

func TestNumIslands(t *testing.T) {
	grid := helper.CharBoardFromJSON(`[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`)
	expected := 1
	assert.Equal(t, expected, solve0200.NumIslands(grid))
}
