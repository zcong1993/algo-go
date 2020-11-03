package solve0106

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 106
@title 从中序与后序遍历序列构造二叉树
@difficulty 中等
@tags tree,depth-first-search,array
@draft false
@link https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
@frontendId 106
*/
// 中序遍历 inorder =&nbsp;[9,3,15,20,7]
// 后序遍历 postorder = [9,15,7,20,3]
/**
	3
   / \
  9  20
    /  \
   15   7
*/
// postorder 最后一个元素为根节点, inorder 通过根节点把树分成左右
// 可以得到左子树 size, 递归构建左右子树
func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	var build func(iStart, iEnd, pStart, pEnd int) *TreeNode
	build = func(iStart, iEnd, pStart, pEnd int) *TreeNode {
		if iStart > iEnd || pStart > pEnd {
			return nil
		}
		rootVal := postorder[pEnd]
		root := &TreeNode{Val: rootVal}
		index := -1
		for i := iStart; i <= iEnd; i++ {
			if inorder[i] == rootVal {
				index = i
				break
			}
		}
		leftSize := index - iStart
		root.Left = build(iStart, index-1, pStart, pStart+leftSize-1)
		root.Right = build(index+1, iEnd, pStart+leftSize, pEnd-1)
		return root
	}
	return build(0, len(inorder)-1, 0, len(postorder)-1)
}
