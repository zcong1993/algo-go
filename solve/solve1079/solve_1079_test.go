package solve1079_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve1079"
)

func TestSumRootToLeaf(t *testing.T) {
	input := tree.Deserialize("1,0,0,#,#,1,#,#,1,0,#,#,1,#,#")
	expected := 22
	assert.Equal(t, expected, solve1079.SumRootToLeaf(input))

	input = tree.Deserialize("1,1,#,#,#")
	expected = 3
	assert.Equal(t, expected, solve1079.SumRootToLeaf(input))
}
