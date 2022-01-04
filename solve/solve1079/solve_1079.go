package solve1079

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 1079
 * @title 从根到叶的二进制数之和
 * @difficulty 简单
 * @tags tree,depth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/
 * @frontendId 1022
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 从根到叶所以需要前序遍历
// 2. 不管是二进制还是十进制, 当前层都等于 prev * 进制 + cur

func SumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	var preorder func(node *TreeNode, prev int)
	preorder = func(node *TreeNode, prev int) {
		if node == nil {
			return
		}

		cur := prev*2 + node.Val

		if node.Left == nil && node.Right == nil {
			sum += cur
		} else {
			preorder(node.Left, cur)
			preorder(node.Right, cur)
		}
	}

	preorder(root, 0)
	return sum
}
