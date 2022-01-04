package solve0102

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 102
 * @title 二叉树的层序遍历
 * @difficulty 中等
 * @tags tree,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 * @frontendId 102
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func LevelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var level []int
		var next []*TreeNode

		for _, cur := range queue {
			level = append(level, cur.Val)

			if cur.Left != nil {
				next = append(next, cur.Left)
			}

			if cur.Right != nil {
				next = append(next, cur.Right)
			}
		}

		res = append(res, level)
		queue = next
	}

	return res
}
