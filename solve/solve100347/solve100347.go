package solve100347

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 100347
@title 二叉树的最近公共祖先
@difficulty 简单
@tags tree
@draft false
@link https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 异常情况
	if root == nil || p == nil || q == nil {
		return root
	}
	// 根节点等于任意一个节点, 这个节点就是最近公共祖先
	if p == root || q == root {
		return root
	}
	// 递归计算出以左右子树为根的最近公共祖先
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	// 左右子树都找到了, 肯定只能是根节点了
	if left != nil && right != nil {
		return root
	}
	// 如果左边没找到, 返回右边结果, 反之亦然
	if left == nil {
		return right
	}
	return left
}
