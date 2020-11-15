package solve100174

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 100174
 * @title 最小高度树
 * @difficulty 简单
 * @tags tree,depth-first-search
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-height-tree-lcci/
 * @frontendId 面试题 04.02
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(l, h int) *TreeNode
	helper = func(l, h int) *TreeNode {
		if l > h {
			return nil
		}
		mid := l + (h-l)/2
		// 有序数组中间值为根
		// 左边构建左子树
		// 右边构建右子树
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(l, mid-1)
		root.Right = helper(mid+1, h)
		return root
	}
	return helper(0, len(nums)-1)
}

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}
