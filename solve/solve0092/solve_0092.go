package solve0092

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 92
 * @title 反转链表 II
 * @difficulty 中等
 * @tags linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/reverse-linked-list-ii/
 * @frontendId 92
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/java-shuang-zhi-zhen-tou-cha-fa-by-mu-yi-cheng-zho

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
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

func ReverseBetween2(head *ListNode, left int, right int) *ListNode {
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

func ReverseBetween3(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left == 1 {
		return ReverseN2(head, right)
	}

	head.Next = ReverseBetween3(head.Next, left-1, right-1)
	return head
}
