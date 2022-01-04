package solve0226

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 226
 * @title 翻转二叉树
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/invert-binary-tree/
 * @frontendId 226
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = root.Right
	root.Right = left

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
