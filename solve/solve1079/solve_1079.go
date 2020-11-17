package solve1079

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 1079
 * @title 从根到叶的二进制数之和
 * @difficulty 简单
 * @tags tree
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

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var preorder func(root *TreeNode, parent int)
	preorder = func(root *TreeNode, parent int) {
		if root == nil {
			return
		}
		// 每到下一层, 将 parent 和左移一位加上当前节点值
		parent = parent<<1 + root.Val
		// 如果是叶节点, 保存结果, 否则继续向下递归
		if root.Left == nil && root.Right == nil {
			sum += parent
		} else {
			preorder(root.Left, parent)
			preorder(root.Right, parent)
		}
	}
	preorder(root, 0)
	return sum
}

func SumRootToLeaf(root *TreeNode) int {
	return sumRootToLeaf(root)
}
