package solve1050

import (
	"sort"

	. "github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0105"
)
import . "github.com/zcong1993/algo-go/pkg/helper"

/**
@index 1050
@title 前序遍历构造二叉搜索树
@difficulty 中等
@tags tree
@draft false
@link https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/
@frontendId 1008
*/
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	var helper func(min, max int) *TreeNode
	// bst 根永远大于左子树所有值, 小于右子树所有值
	// 通过根能够限制左右子树节点范围, 递归构建左右子树
	helper = func(min, max int) *TreeNode {
		// index out of range, return
		if index >= len(preorder) {
			return nil
		}
		val := preorder[index]
		// 如果节点不在该范围, 则说明这个节点应该在另一边, return nil
		if val < min || val > max {
			return nil
		}
		// 当前节点在范围内, 作为跟节点
		root := &TreeNode{Val: val}
		// 用掉一个值, 指针右移
		index++
		// 递归构建左子树, min < 左子树值 < 当前节点值
		root.Left = helper(min, val)
		// // 递归构建右子树, 当前节点值 < 右子树值 < max
		root.Right = helper(val, max)
		return root
	}
	// 从 -inf +inf 开始构建
	root := helper(MinInt, MaxInt)
	return root
}

func BstFromPreorder(preorder []int) *TreeNode {
	return bstFromPreorder(preorder)
}

// bst inorder 为单调递增, sort preorder 可以得到 inorder
// 然后变成了 105 题目, 根据前序中序构建树
func BstFromPreorder2(preorder []int) *TreeNode {
	inorder := make([]int, len(preorder))
	copy(inorder, preorder)
	sort.Ints(inorder)
	return solve0105.BuildTree(preorder, inorder)
}
