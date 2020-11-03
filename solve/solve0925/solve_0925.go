package solve0925

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 925
@title 根据前序和后序遍历构造二叉树
@difficulty 中等
@tags tree
@draft false
@link https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
@frontendId 889
*/
// pre = [1,2,4,5,3,6,7]
// post = [4,5,2,6,7,3,1]
// 前序遍历第一个元素是根节点, 第二个元素是左子树的根节点
// 通过左子树根节点结合后续遍历, 可以知道左右子树的元素数目
// 递归构建左右子树
func ConstructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 || len(pre) != len(post) {
		return nil
	}
	var build func(prStart, prEnd, poStart, poEnd int) *TreeNode
	build = func(prStart, prEnd, poStart, poEnd int) *TreeNode {
		if prStart > prEnd || poStart > poEnd {
			return nil
		}
		rootVal := pre[prStart]
		root := &TreeNode{Val: rootVal}
		nextRootPrIndex := prStart + 1
		if nextRootPrIndex > prEnd {
			return root
		}
		nextRootVal := pre[nextRootPrIndex]
		index := -1
		for i := poStart; i <= poEnd; i++ {
			if post[i] == nextRootVal {
				index = i
				break
			}
		}
		leftSize := index - poStart + 1
		root.Left = build(prStart+1, prStart+leftSize, poStart, index)
		root.Right = build(prStart+leftSize+1, prEnd, index+1, poEnd-1)
		return root
	}
	return build(0, len(pre)-1, 0, len(post)-1)
}
