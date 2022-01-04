package solve0572

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 572
 * @title 另一棵树的子树
 * @difficulty 简单
 * @tags tree,depth-first-search,binary-tree,string-matching,hash-function
 * @draft false
 * @link https://leetcode-cn.com/problems/subtree-of-another-tree/
 * @frontendId 572
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSame(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}

	return isSame(node1.Left, node2.Left) && isSame(node1.Right, node2.Right)
}

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	return isSame(root, subRoot) || IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
