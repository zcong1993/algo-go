package solve0572_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0572"
)

func TestIsSubtree(t *testing.T) {
	tree1 := tree.Deserialize("3,4,1,#,#,2,#,#,5,#,#")
	tree2 := tree.Deserialize("4,1,#,#,2,#,#")
	tree3 := tree.Deserialize("3,4,1,#,#,2,0,#,#,#,5,#,#")

	assert.True(t, solve0572.IsSubtree(tree1, tree2))
	assert.False(t, solve0572.IsSubtree(tree3, tree2))
}
