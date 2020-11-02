package solve0463

import (
	"testing"

	"github.com/zcong1993/algo-go/pkg/helper"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	grid := helper.GridFromJSON("[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]")
	assert.Equal(t, 16, IslandPerimeter(grid))
}
