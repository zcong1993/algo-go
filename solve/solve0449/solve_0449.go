package solve0449

import (
	"strconv"
	"strings"

	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
 * @index 449
 * @title 序列化和反序列化二叉搜索树
 * @difficulty 中等
 * @tags tree,depth-first-search,breadth-first-search,design,binary-search-tree,string,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/serialize-and-deserialize-bst/
 * @frontendId 449
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
	empty = "#"
	step  = ","
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func mustAtoi(s string) int {
	in, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return in
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	res := strings.Builder{}

	if root == nil {
		return empty
	}

	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			res.WriteString(empty)
			res.WriteString(step)
			return
		}

		res.WriteString(strconv.Itoa(node.Val))
		res.WriteString(step)

		pre(node.Left)
		pre(node.Right)
	}

	pre(root)

	return strings.TrimSuffix(res.String(), step)
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	if len(data) == 0 || data == empty {
		return nil
	}

	arr := strings.Split(data, step)
	i := 0

	var build func() *TreeNode
	build = func() *TreeNode {
		if i >= len(arr) {
			return nil
		}
		last := arr[i]
		i++

		if last == empty {
			return nil
		}

		root := &TreeNode{Val: mustAtoi(last)}
		root.Left = build()
		root.Right = build()

		return root
	}

	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
