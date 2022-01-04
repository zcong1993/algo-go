package solve1000019

import (
	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 1000019
 * @title BiNode
 * @difficulty 简单
 * @tags stack,tree,depth-first-search,binary-search-tree,linked-list,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binode-lcci/
 * @frontendId 面试题 17.12
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 题目要求顺序为中序遍历形成的链表
// 2. 中序遍历, 拼接链表顺序为 左 -> root -> 右
// 3. 需要考虑 root 有没有左节点两种情况

func ConvertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head := root

	if root.Left != nil {
		left := ConvertBiNode(root.Left)
		leftEnd := left
		for leftEnd.Right != nil {
			leftEnd = leftEnd.Right
		}
		leftEnd.Right = root
		head = left
	}

	root.Left = nil
	root.Right = ConvertBiNode(root.Right)
	return head
}
