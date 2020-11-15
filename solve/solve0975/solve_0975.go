package solve0975

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
 * @index 975
 * @title 二叉搜索树的范围和
 * @difficulty 简单
 * @tags tree,depth-first-search,recursion
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 当前节点在范围内, 更新结果
		// 递归搜索左右子树
		if root.Val >= low && root.Val <= high {
			sum += root.Val
			helper(root.Left)
			helper(root.Right)
		} else {
			// 节点不在范围内, 由于是搜索树
			// 所以可以排除一边子树
			// 当前节点值小于最小搜索值时, 搜索右子树就可以了
			if root.Val < low {
				helper(root.Right)
				// 当前节点值大于最大搜索值时, 搜索左子树
			} else if root.Val > high {
				helper(root.Left)
			}
		}
	}
	helper(root)
	return sum
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	return rangeSumBST(root, low, high)
}
