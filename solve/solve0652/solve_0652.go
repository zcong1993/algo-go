package solve0652

import (
	"fmt"

	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 652
 * @title 寻找重复的子树
 * @difficulty 中等
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/find-duplicate-subtrees/
 * @frontendId 652
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 后序遍历才知道子树结构
// 2. 使用 map 记录是否有重复
// 3. 注意结果去重

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}

	rMap := make(map[*TreeNode]struct{})
	strMap := make(map[string]bool)

	var post func(node *TreeNode) string
	post = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := post(node.Left)
		right := post(node.Right)
		// 左父右会出问题
		str := fmt.Sprintf("%d,%s,%s", node.Val, left, right)

		if added, ok := strMap[str]; ok {
			if !added {
				rMap[node] = struct{}{}
				strMap[str] = true
			}
		} else {
			strMap[str] = false
		}

		return str
	}

	post(root)

	res := make([]*TreeNode, len(rMap))
	i := 0
	for n := range rMap {
		res[i] = n
		i++
	}

	return res
}
