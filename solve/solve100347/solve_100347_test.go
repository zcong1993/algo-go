package solve100347_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve100347"
)

func TestLowestCommonAncestor(t *testing.T) {
	input := tree.Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#")
	res1 := solve100347.LowestCommonAncestor(input, &tree.TreeNode{Val: 5}, &tree.TreeNode{Val: 1})
	res2 := solve100347.LowestCommonAncestor(input, &tree.TreeNode{Val: 5}, &tree.TreeNode{Val: 4})

	assert.Equal(t, 3, res1.Val)
	assert.Equal(t, 5, res2.Val)
}
