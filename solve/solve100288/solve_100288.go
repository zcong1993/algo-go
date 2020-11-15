package solve100288

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 100288
 * @title 二叉树的镜像
 * @difficulty 简单
 * @tags tree
 * @draft false
 * @link https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
 * @frontendId 剑指 Offer 27
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左右交换左右子节点
		root.Left, root.Right = root.Right, root.Left
		// 递归操作左右子树
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return root
}

func MirrorTree(root *TreeNode) *TreeNode {
	return mirrorTree(root)
}
