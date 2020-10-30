package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	assert.True(t, IsSameTree(nil, nil))
	assert.False(t, IsSameTree(&TreeNode{Val: 0}, nil))
	assert.False(t, IsSameTree(nil, &TreeNode{Val: 0}))
	assert.False(t, IsSameTree(&TreeNode{Val: 0}, &TreeNode{Val: 1}))
	assert.True(t, IsSameTree(&TreeNode{Val: 1}, &TreeNode{Val: 1}))

	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}

	root1 := &TreeNode{Val: 0}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	assert.True(t, IsSameTree(root, root1))
}

func TestSerialize(t *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	assert.Equal(t, "0,1,3,#,#,#,2,#,#", Serialize(root))
}

func TestDeserialize(t *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	str := Serialize(root)
	deTree := Deserialize(str)
	assert.True(t, IsSameTree(root, deTree))
}

func TestTraversal(t *testing.T) {
	root := Deserialize("0,1,3,#,#,#,2,4,#,#,#")
	assert.Equal(t, []int{0, 1, 3, 2, 4}, Traversal(root, PREORDER))
	assert.Equal(t, []int{3, 1, 0, 4, 2}, Traversal(root, INORDER))
	assert.Equal(t, []int{3, 1, 4, 2, 0}, Traversal(root, POSTORDER))
}
