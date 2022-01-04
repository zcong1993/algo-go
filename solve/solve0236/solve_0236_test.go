package solve0236_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0236"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := tree.Deserialize("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#")
	var node1 *tree.TreeNode
	var node3 *tree.TreeNode
	var node4 *tree.TreeNode
	var node5 *tree.TreeNode

	tree.TraversalWithFn(root, tree.INORDER, func(node *tree.TreeNode) {
		if node.Val == 1 {
			node1 = node
		} else if node.Val == 3 {
			node3 = node
		} else if node.Val == 4 {
			node4 = node
		} else if node.Val == 5 {
			node5 = node
		}
	})

	res1 := solve0236.LowestCommonAncestor(root, node5, node1)
	res2 := solve0236.LowestCommonAncestor(root, node5, node4)

	assert.Equal(t, node3, res1)
	assert.Equal(t, node5, res2)
}
