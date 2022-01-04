package solve0102_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0102"
)

func TestLevelOrder(t *testing.T) {
	input := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	expected := helper.GridFromJSON(`[
  [3],
  [9,20],
  [15,7]
]`)

	assert.Equal(t, expected, solve0102.LevelOrder(input))
}
