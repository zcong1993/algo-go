package solve0230

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 230
 * @title 二叉搜索树中第K小的元素
 * @difficulty 中等
 * @tags tree,depth-first-search,binary-search-tree,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
 * @frontendId 230
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 中序有序
// 2. 找到后剪枝掉剩下节点

func KthSmallest(root *TreeNode, k int) int {
	if root == nil || k < 0 {
		return 0
	}

	i := 0
	var res int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		i++
		if i == k {
			res = node.Val
			return
		}

		if i < k {
			inorder(node.Right)
		}
	}

	inorder(root)

	return res
}
