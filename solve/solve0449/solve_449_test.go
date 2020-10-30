package solve0449

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestSerialize(t *testing.T) {
	c := Constructor()
	root := &tree.TreeNode{Val: 0}
	root.Left = &tree.TreeNode{Val: 1}
	root.Right = &tree.TreeNode{Val: 2}
	root.Left.Left = &tree.TreeNode{Val: 3}
	assert.Equal(t, "0,1,3,#,#,#,2,#,#", c.serialize(root))
}

func TestDeserialize(t *testing.T) {
	c := Constructor()
	root := &tree.TreeNode{Val: 0}
	root.Left = &tree.TreeNode{Val: 1}
	root.Right = &tree.TreeNode{Val: 2}
	root.Left.Left = &tree.TreeNode{Val: 3}
	root.Right.Left = &tree.TreeNode{Val: 4}
	str := c.serialize(root)
	deTree := c.deserialize(str)
	assert.True(t, tree.IsSameTree(root, deTree))
}
