package solve0102

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 102
@title 二叉树的层序遍历
@difficulty 中等
@tags tree,breadth-first-search
@draft false
@link https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
*/
func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		next := make([]*TreeNode, 0)
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, level)
		queue = next
	}
	return res
}
