package solve0101

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 101
 * @title 对称二叉树
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/symmetric-tree/
 * @frontendId 101
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
}

func IsSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}
