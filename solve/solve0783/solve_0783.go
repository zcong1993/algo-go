package solve0783

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 783
 * @title 二叉搜索树中的搜索
 * @difficulty 简单
 * @tags tree
 * @draft false
 * @link https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
 * @frontendId 700
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	return searchBST(root, val)
}
