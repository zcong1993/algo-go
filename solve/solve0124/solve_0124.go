package solve0124

import (
	"math"

	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 124
 * @title 二叉树中的最大路径和
 * @difficulty 困难
 * @tags tree,depth-first-search,dynamic-programming,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
 * @frontendId 124
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }.
 */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MinInt32
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 只有大于 0 才会被选取
		left := max(postOrder(root.Left), 0)
		right := max(postOrder(root.Right), 0)
		// 更新结果
		sum := left + right + root.Val
		res = max(res, sum)
		// 不选取当前节点作为最高点, 子树只能选取大的一半
		return root.Val + max(left, right)
	}

	postOrder(root)
	return res
}
