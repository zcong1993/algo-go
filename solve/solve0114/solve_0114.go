package solve0114

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 114
 * @title 二叉树展开为链表
 * @difficulty 中等
 * @tags stack,tree,depth-first-search,linked-list,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
 * @frontendId 114
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 先将左右子树展开
	Flatten(root.Left)
	Flatten(root.Right)

	// 将左子树作为右节点
	right := root.Right
	root.Right = root.Left
	// 左节点置空
	root.Left = nil

	// 将右子树拼在树尾部(左节点尾)
	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}
