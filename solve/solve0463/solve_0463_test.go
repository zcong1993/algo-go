package solve0463

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo: common util
func gridFromJSON(input string) [][]int {
	var grid [][]int
	_ = json.Unmarshal([]byte(input), &grid)
	return grid
}

func TestIslandPerimeter(t *testing.T) {
	grid := gridFromJSON("[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]")
	assert.Equal(t, 16, IslandPerimeter(grid))
}
