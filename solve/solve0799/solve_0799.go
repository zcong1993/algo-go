package solve0799

import (
	"math"

	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 799
 * @title 二叉搜索树节点最小距离
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-search-tree,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
 * @frontendId 783
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min2(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

// 1. BST 中序遍历有序
// 2. 差值最小的两个节点肯定是相邻的

func MinDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min := math.MaxInt64
	var preNode *TreeNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if preNode != nil {
			min = min2(min, node.Val-preNode.Val)
		}
		preNode = node

		inorder(node.Right)
	}

	inorder(root)

	return min
}
