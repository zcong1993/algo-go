package solve0105

import . "github.com/zcong1993/algo-go/pkg/tree"

/**
@index 105
@title 从前序与中序遍历序列构造二叉树
@difficulty 中等
@tags tree,depth-first-search,array
@draft false
@link https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/

// preorder [1,2,4,5,3,6,7]
// inorder [4,2,5,1,6,3,7]
// preorder 顺序为 根, 左子树, 右子树
// inorder 顺序为 左子树, 根, 右子树
// 先从 preorder 拿到根, 查询 inorder 就知道左子树的节点数, 就能分出来左右子树向下递归
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(inorder) != len(preorder) {
		return nil
	}
	var build func(pStart, pEnd, iStart, iEnd int) *TreeNode
	build = func(pStart, pEnd, iStart, iEnd int) *TreeNode {
		if pStart > pEnd || iStart > iEnd {
			return nil
		}
		rootVal := preorder[pStart]
		root := &TreeNode{Val: rootVal}
		index := -1
		for i, val := range inorder {
			if val == rootVal {
				index = i
				break
			}
		}
		leftSize := index - iStart
		root.Left = build(pStart+1, pStart+leftSize, iStart, index-1)
		root.Right = build(pStart+leftSize+1, pEnd, index+1, iEnd)
		return root
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
