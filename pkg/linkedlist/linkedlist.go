package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromArray(input []int) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for _, v := range input {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func ToArray(node *ListNode) []int {
	res := make([]int, 0)
	n := node
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	return res
}

func Reverse(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	var prev *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func Reverse2(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	last := Reverse2(node.Next)
	node.Next.Next = node
	node.Next = nil
	return last
}

func Combine(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = node1
	end := dummy
	for end.Next != nil {
		end = end.Next
	}
	end.Next = node2
	return dummy.Next
}

func ReverseN(node *ListNode, n int) *ListNode {
	if node == nil {
		return nil
	}

	var newHead *ListNode
	var prev *ListNode
	cur := node

	// 注意前 n 个
	for i := 0; i < n; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		// 最后一个节点翻转后
		// 将初始的头连到下一个节点
		// 注意此时的 cur 已经指向下一个节点了
		if i == n-1 {
			newHead = prev
			node.Next = cur
		}
	}

	return newHead
}

func ReverseN2(head *ListNode, n int) *ListNode {
	var successor *ListNode
	var reverseN2 func(node *ListNode, n int) *ListNode
	reverseN2 = func(node *ListNode, n int) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}
		if n == 1 {
			successor = node.Next
			return node
		}
		last := reverseN2(node.Next, n-1)
		node.Next.Next = node
		node.Next = successor
		return last
	}

	return reverseN2(head, n)
}

func ReverseNM(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 虚节点方便处理翻转全部的状态
	dummy := &ListNode{Val: 0}
	dummy.Next = head

	var (
		prev *ListNode
		tail *ListNode // tail 保存开始翻转的第一个节点, 后续需要将它的 next 连接到 right+1 节点
		g    = dummy   // g 保存翻转开始前一个节点 left-1, 后续需要将它的 next 连接到 right 节点
		cur  = dummy.Next
	)

	// 先将 g cur 移动到开始位置
	for i := 0; i < left-1; i++ {
		g = g.Next
		cur = cur.Next
	}

	// 翻转链表
	for i := 0; i < right-left+1; i++ {
		// left 节点记录起来
		if i == 0 {
			tail = cur
		}
		// 翻转当前节点
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		// 到 right 节点时, 添加 g 和 tail 节点关系
		if i == right-left {
			g.Next = prev
			tail.Next = cur
		}
	}

	return dummy.Next
}

func ReverseNM2(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head

	g, p := dummy, dummy.Next

	for i := 0; i < left-1; i++ {
		g = g.Next
		p = p.Next
	}

	for i := 0; i < right-left; i++ {
		removed := p.Next
		p.Next = p.Next.Next

		removed.Next = g.Next
		g.Next = removed
	}

	return dummy.Next
}

func ReverseNM3(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left == 1 {
		return ReverseN2(head, right)
	}

	head.Next = ReverseNM3(head.Next, left-1, right-1)
	return head
}

func ReverseNM4(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	g, p := dummy, dummy.Next

	for i := 0; i < left-1; i++ {
		g, p = g.Next, p.Next
	}

	for i := 0; i < right-left; i++ {
		removed := p.Next
		p.Next = p.Next.Next

		removed.Next = g.Next
		g.Next = removed
	}

	return dummy.Next
}
