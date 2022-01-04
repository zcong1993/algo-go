package solve0925

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 925
 * @title 根据前序和后序遍历构造二叉树
 * @difficulty 中等
 * @tags tree,array,hash-table,divide-and-conquer,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
 * @frontendId 889
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
// 2. 前序 root 下一个节点就是左子树的 root
// 3. 根据左子树 root, 搜索后序数组得到左子树长度
// 4. 分割左右子树递归

func ConstructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 || len(preorder) != len(postorder) {
		return nil
	}

	l := len(preorder)

	var build func(prStart, prEnd, poStart, poEnd int) *TreeNode
	build = func(prStart, prEnd, poStart, poEnd int) *TreeNode {
		if prStart > prEnd || poStart > poEnd {
			return nil
		}

		rootVal := preorder[prStart]
		root := &TreeNode{Val: rootVal}

		if prStart == prEnd {
			return root
		}

		nextRootIdx := prStart + 1
		nextRootVal := preorder[nextRootIdx]
		postIdx := -1

		for i := poStart; i <= poEnd; i++ {
			if nextRootVal == postorder[i] {
				postIdx = i
				break
			}
		}

		leftLen := postIdx - poStart + 1

		root.Left = build(prStart+1, prStart+leftLen, poStart, postIdx)
		root.Right = build(prStart+leftLen+1, prEnd, postIdx+1, poEnd)
		return root
	}

	return build(0, l-1, 0, l-1)
}
