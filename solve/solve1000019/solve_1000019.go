package solve1000019

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 1000019
@title BiNode
@difficulty 简单
@tags tree,binary-search-tree,recursion
@draft false
@link https://leetcode-cn.com/problems/binode-lcci/
@frontendId 面试题 17.12
*/
func ConvertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// prev 节点为上转换好的链表的尾节点
	var head, prev *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 将左边转换好
		inorder(root.Left)
		// 最左端节点, 保存为头
		// 更新 prev 节点
		if head == nil {
			head = root
			prev = root
		} else {
			// 左边转换好的尾节点右子节点指向当前节点
			prev.Right = root
			// 更新 prev 为当前节点
			prev = root
		}
		// 删除当前节点左节点
		root.Left = nil
		// 递归处理剩下的节点
		inorder(root.Right)
	}
	inorder(root)
	// 返回头
	return head
}
