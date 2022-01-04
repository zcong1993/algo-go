package solve0111

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 111
 * @title 二叉树的最小深度
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
 * @frontendId 111
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 唯一需要注意的是只有半边的情况

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	} else if root.Right == nil {
		return MinDepth(root.Left) + 1
	}

	minLeft := MinDepth(root.Left)
	minRight := MinDepth(root.Right)

	if minLeft < minRight {
		return 1 + minLeft
	} else {
		return 1 + minRight
	}
}
