package solve0111

import (
	. "github.com/zcong1993/algo-go/pkg/helper"
	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 111
 * @title 二叉树的最小深度
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	res := 0
	if root.Left == nil {
		res = minDepth(root.Right)
	} else if root.Right == nil {
		res = minDepth(root.Left)
	} else {
		res = Min2(minDepth(root.Left), minDepth(root.Right))
	}
	return res + 1
}

func MinDepth(root *TreeNode) int {
	return minDepth(root)
}
