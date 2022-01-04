package solve0103

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 103
 * @title 二叉树的锯齿形层序遍历
 * @difficulty 中等
 * @tags tree,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
 * @frontendId 103
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	flag := 1
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

		if flag < 0 {
			Reverse(level)
		}

		res = append(res, level)
		flag *= -1

		queue = next
	}
	return res
}
