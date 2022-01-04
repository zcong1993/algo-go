package solve0105

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 105
 * @title 从前序与中序遍历序列构造二叉树
 * @difficulty 中等
 * @tags tree,array,hash-table,divide-and-conquer,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 * @frontendId 105
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 前序遍历 root 节点在头, 后序在尾
// 2. 中序遍历 root 左右就是左右子树
// 3. 根据 root 从中序遍历得到左子树长度
// 4. 分割左右子树递归

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}

	l := len(preorder)

	var build func(prStart, prEnd, inStart, inEnd int) *TreeNode
	build = func(prStart, prEnd, inStart, inEnd int) *TreeNode {
		if prStart > prEnd || inStart > inEnd {
			return nil
		}

		rootVal := preorder[prStart]
		root := &TreeNode{Val: rootVal}

		inIdx := -1
		for i := inStart; i <= inEnd; i++ {
			if inorder[i] == rootVal {
				inIdx = i
				break
			}
		}

		leftLen := inIdx - inStart

		root.Left = build(prStart+1, prStart+leftLen, inStart, inIdx-1)
		root.Right = build(prStart+leftLen+1, prEnd, inIdx+1, inEnd)

		return root
	}

	return build(0, l-1, 0, l-1)
}
