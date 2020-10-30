package solve0563

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 563
@title 二叉树的坡度
@difficulty 简单
@tags tree
@draft false
@link https://leetcode-cn.com/problems/binary-tree-tilt
*/
func FindTilt(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}

	// 后序遍历树, 函数返回子树节点和
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := helper(root.Left), helper(root.Right)
		// 结果加上该子树坡度
		res += abs(left - right)
		// 子树节点和 = 左子树 + 右子树 + 根节点
		return left + right + root.Val
	}
	helper(root)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
