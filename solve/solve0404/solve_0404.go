package solve0404

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 404
 * @title 左叶子之和
 * @difficulty 简单
 * @tags tree
 * @draft false
 * @link https://leetcode-cn.com/problems/sum-of-left-leaves/
 * @frontendId 404
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var post func(root *TreeNode)
	post = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		}
		post(root.Left)
		post(root.Right)
	}
	post(root)
	return sum
}

func SumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeaves(root)
}
