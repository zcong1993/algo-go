package solve0116

/**
 * @index 116
 * @title 填充每个节点的下一个右侧节点指针
 * @difficulty 中等
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
 * @frontendId 116
 */

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Left *Node
*     Right *Node
*     Next *Node
* }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connect2(root.Left, root.Right)
	return root
}

func connect2(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	connect2(left.Left, left.Right)
	connect2(right.Left, right.Right)
	connect2(left.Right, right.Left)
}

// 从右向左层序遍历, 将上一个节点作为当前节点 next

func Connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		var next []*Node
		var pre *Node

		for _, node := range queue {
			node.Next = pre
			pre = node
			if node.Right != nil {
				next = append(next, node.Right)
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
		}

		queue = next
	}

	return root
}
