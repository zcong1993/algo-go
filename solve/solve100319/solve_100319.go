package solve100319

import (
	. "github.com/zcong1993/algo-go/pkg/helper"
	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 100319
 * @title 二叉树的深度
 * @difficulty 简单
 * @tags tree,depth-first-search
 * @draft false
 * @link https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
 * @frontendId 剑指 Offer 55 - I
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }.
 */
func maxDepth(root *TreeNode) int {
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := postOrder(root.Left), postOrder(root.Right)
		// 树高 = 1(当前层) + 左右子树树深大的
		return 1 + Max2(l, r)
	}
	return postOrder(root)
}

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}
