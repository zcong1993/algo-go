package solve100347

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &tree.TreeNode{Val: 3}
	node5 := &tree.TreeNode{Val: 5}
	node1 := &tree.TreeNode{Val: 1}
	node6 := &tree.TreeNode{Val: 6}
	node2 := &tree.TreeNode{Val: 2}

	root.Left = node5
	root.Right = node1
	node5.Left = node6
	node5.Right = node2
	assert.Equal(t, root, LowestCommonAncestor(root, node5, node1))
	assert.Equal(t, node5, LowestCommonAncestor(root, node6, node2))
	assert.Equal(t, root, LowestCommonAncestor(root, node6, node1))
}
