package solve0106

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 106
 * @title 从中序与后序遍历序列构造二叉树
 * @difficulty 中等
 * @tags tree,array,hash-table,divide-and-conquer,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
 * @frontendId 106
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}

	var build func(inStart, inEnd, poStart, poEnd int) *TreeNode
	build = func(inStart, inEnd, poStart, poEnd int) *TreeNode {
		if inStart > inEnd || poStart > poEnd {
			return nil
		}

		rootVal := postorder[poEnd]
		root := &TreeNode{Val: rootVal}

		inIdx := -1
		for i := inStart; i <= inEnd; i++ {
			if inorder[i] == rootVal {
				inIdx = i
				break
			}
		}

		leftLen := inIdx - inStart

		root.Left = build(inStart, inIdx-1, poStart, poStart+leftLen-1)
		root.Right = build(inIdx+1, inEnd, poStart+leftLen, poEnd-1)

		return root
	}

	l := len(inorder)

	return build(0, l-1, 0, l-1)
}
