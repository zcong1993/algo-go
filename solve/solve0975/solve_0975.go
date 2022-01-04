package solve0975

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 975
 * @title 二叉搜索树的范围和
 * @difficulty 简单
 * @tags tree,depth-first-search,binary-search-tree,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/range-sum-of-bst/
 * @frontendId 938
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. BST 中序有序
// 2. 注意剪枝大于 high 的节点

func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil || low > high {
		return 0
	}

	sum := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}

		if node.Val <= high {
			inorder(node.Right)
		}
	}

	inorder(root)
	return sum
}

// 更好利用 BST 性质, 中序遍历, 更多剪枝
// 可以直接放弃 < low || > high 的节点

func RangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		cur := node.Val

		if cur >= low && cur <= high {
			sum += cur
			preorder(node.Left)
			preorder(node.Right)
		} else if cur < low {
			preorder(node.Right)
		} else if cur > high {
			preorder(node.Left)
		}
	}

	preorder(root)
	return sum
}
