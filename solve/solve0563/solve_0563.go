package solve0563

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 563
 * @title 二叉树的坡度
 * @difficulty 简单
 * @tags tree,depth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-tilt/
 * @frontendId 563
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func FindTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	var sum func(node *TreeNode) int
	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 后序遍历求和
		leftSum := sum(node.Left)
		rightSum := sum(node.Right)
		// 得到左右子树和求梯度
		res += abs(leftSum - rightSum)

		return leftSum + rightSum + node.Val
	}

	sum(root)

	return res
}
