package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Order int

const (
	PREORDER Order = iota + 1
	INORDER
	POSTORDER
)

const (
	EMPTY = "#"
	STEP  = ","
)

// 遍历树
func Traversal(root *TreeNode, order Order) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var t func(root *TreeNode)
	t = func(root *TreeNode) {
		if root == nil {
			return
		}
		if order == PREORDER {
			res = append(res, root.Val)
		}
		t(root.Left)
		if order == INORDER {
			res = append(res, root.Val)
		}
		t(root.Right)
		if order == POSTORDER {
			res = append(res, root.Val)
		}
	}
	t(root)
	return res
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

// 序列化
func Serialize(root *TreeNode) string {
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

// 反序列化
func Deserialize(str string) *TreeNode {
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

// 判断两棵树是否相同
func IsSameTree(root1, root2 *TreeNode) bool {
	// 两树都为空, true
	if root1 == nil && root2 == nil {
		return true
	}
	// 一个节点为空, 或者两个节点值不相同, false
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	// 递归对比左子树和右子树
	return IsSameTree(root1.Left, root2.Left) && IsSameTree(root1.Right, root2.Right)
}
