package solve0617

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 617
 * @title 合并二叉树
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
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

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	// 原地修改
	root := root1
	root.Val += root2.Val

	root.Left = MergeTrees(root1.Left, root2.Left)
	root.Right = MergeTrees(root1.Right, root2.Right)

	return root
}
