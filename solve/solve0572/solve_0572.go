package solve0572

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 572
@title 另一个树的子树
@difficulty 简单
@tags tree
@draft false
@link https://leetcode-cn.com/problems/subtree-of-another-tree
*/
func IsSubtree(s *TreeNode, t *TreeNode) bool {
	// 两树都为空, true
	if s == nil && t == nil {
		return true
	}
	// 任意一个为空, false
	if s == nil || t == nil {
		return false
	}
	// 子树 = 两树相同 || t 为 s 左子树的子树 || t 为 s 右子树的子树
	return IsSameTree(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}
