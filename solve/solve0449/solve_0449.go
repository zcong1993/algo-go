package solve0449

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/zcong1993/algo-go/pkg/tree"
)

/**
@index 449
@title 序列化和反序列化二叉搜索树
@difficulty 中等
@tags tree
@draft false
@link https://leetcode-cn.com/problems/serialize-and-deserialize-bst/
*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return EMPTY
	}
	sb := strings.Builder{}
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			sb.WriteString(EMPTY)
			sb.WriteString(STEP)
			return
		}
		sb.WriteString(int2string(root.Val))
		sb.WriteString(STEP)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return strings.TrimSuffix(sb.String(), STEP)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(str string) *TreeNode {
	if str == "" || str == EMPTY {
		return nil
	}
	nodes := strings.Split(str, STEP)
	index := 0
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if index > len(nodes)-1 {
			return nil
		}
		last := nodes[index]
		index++
		if last == EMPTY {
			return nil
		}
		root := &TreeNode{Val: mustAtoi(last)}
		root.Left = helper()
		root.Right = helper()
		return root
	}
	return helper()
}

func int2string(i int) string {
	return fmt.Sprintf("%d", i)
}

func mustAtoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
