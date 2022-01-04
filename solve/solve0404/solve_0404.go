package solve0404

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 404
 * @title 左叶子之和
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
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

// 1. 增加参数标记是否为左节点
// 2. 叶子节点没有任何子节点
// 3. 根不算左节点

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	var helper func(node *TreeNode, isLeft bool)
	helper = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		if isLeft && node.Left == nil && node.Right == nil {
			sum += node.Val
		}

		helper(node.Left, true)
		helper(node.Right, false)
	}

	helper(root, false)

	return sum
}
