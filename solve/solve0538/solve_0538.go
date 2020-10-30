package solve0538

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 538
@title 把二叉搜索树转换为累加树
@difficulty 中等
@tags tree
@draft false
@link https://leetcode-cn.com/problems/convert-bst-to-greater-tree
*/
func ConvertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	var inorder func(root *TreeNode)
	// 先右后左中序遍历, 单调递减
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		// 更新大于等于当前节点值的和
		sum += root.Val
		// 更新当前节点值
		root.Val = sum
		inorder(root.Left)
	}
	inorder(root)
	return root
}
