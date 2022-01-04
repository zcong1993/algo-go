package solve2035_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve2035"
)

func TestCountSubIslands(t *testing.T) {
	grid1 := helper.GridFromJSON(`[[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]]`)
	grid2 := helper.GridFromJSON(`[[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]`)
	expected := 2
	assert.Equal(t, expected, solve2035.CountSubIslands(grid1, grid2))
}
