package solve0236

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 236
 * @title 二叉树的最近公共祖先
 * @difficulty 中等
 * @tags tree,depth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 * @frontendId 236
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}
