package solve1079_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve1079"
)

func TestSumRootToLeaf(t *testing.T) {
	root := tree.Deserialize("1,0,0,#,#,1,#,#,1,0,#,#,1,#,#")
	assert.Equal(t, 22, solve1079.SumRootToLeaf(root))
}
