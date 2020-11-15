package solve0617

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 617
 * @title 合并二叉树
 * @difficulty 简单
 * @tags tree
 * @draft false
 * @link https://leetcode-cn.com/problems/merge-two-binary-trees/
 * @frontendId 617
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	// 一边为空时, 返回不为空的那边
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	// 合并节点值
	root := &TreeNode{Val: t1.Val + t2.Val}
	// 递归合并左右子树
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return mergeTrees(t1, t2)
}
