package solve0538

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 538
 * @title 把二叉搜索树转换为累加树
 * @difficulty 中等
 * @tags tree,depth-first-search,binary-search-tree,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
 * @frontendId 538
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func ConvertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历, 先右后左, 遍历顺序就是从大到小
		// 所以 sum 就是比当前节点大的所有节点和
		inorder(node.Right)
		sum += node.Val
		node.Val = sum
		inorder(node.Left)
	}

	inorder(root)

	return root
}
