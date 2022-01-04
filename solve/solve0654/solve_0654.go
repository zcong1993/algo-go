package solve0654

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 654
 * @title 最大二叉树
 * @difficulty 中等
 * @tags stack,tree,array,divide-and-conquer,binary-tree,monotonic-stack
 * @draft false
 * @link https://leetcode-cn.com/problems/maximum-binary-tree/
 * @frontendId 654
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var build func(start, end int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		max := -1
		var idx int
		for i := start; i <= end; i++ {
			if nums[i] > max {
				max = nums[i]
				idx = i
			}
		}

		root := &TreeNode{Val: max}
		root.Left = build(start, idx-1)
		root.Right = build(idx+1, end)
		return root
	}

	return build(0, len(nums)-1)
}
