package solve1380_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve1380"
)

func TestClosedIsland(t *testing.T) {
	inputs := helper.GridFromJSON(`[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]`)
	expected := 2
	assert.Equal(t, expected, solve1380.ClosedIsland(inputs))
}
