package solve0463

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestIslandPerimeter(t *testing.T) {
	grid := helper.GridFromJSON("[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]")
	assert.Equal(t, 16, IslandPerimeter(grid))
	assert.Equal(t, 0, IslandPerimeter(make([][]int, 0)))
	assert.Equal(t, 0, IslandPerimeter([][]int{{}}))
}
