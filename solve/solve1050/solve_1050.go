package solve1050

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 1050
 * @title 前序遍历构造二叉搜索树
 * @difficulty 中等
 * @tags stack,tree,binary-search-tree,array,binary-tree,monotonic-stack
 * @draft false
 * @link https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/
 * @frontendId 1008
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. bst 前序遍历第一个元素为 root
// 2. 比 root 小的是左子树
// 3. 根据特性分割子树递归构建

func BstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	l := len(preorder)

	var build func(start, end int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		rootVal := preorder[start]
		root := &TreeNode{Val: rootVal}

		if start == end {
			return root
		}

		leftLen := 0

		for i := start + 1; i <= end; i++ {
			if preorder[i] < rootVal {
				leftLen++
			} else {
				break
			}
		}

		root.Left = build(start+1, start+leftLen)
		root.Right = build(start+leftLen+1, end)
		return root
	}

	return build(0, l-1)
}
