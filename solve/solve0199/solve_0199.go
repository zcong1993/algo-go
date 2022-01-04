package solve0199

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 199
 * @title 二叉树的右视图
 * @difficulty 中等
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-right-side-view/
 * @frontendId 199
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func RightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var next []*TreeNode
		var last *TreeNode
		for _, node := range queue {
			last = node
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, last.Val)
		queue = next
	}

	return res
}
