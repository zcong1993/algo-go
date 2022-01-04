package solve0019

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 19
 * @title 删除链表的倒数第 N 个结点
 * @difficulty 中等
 * @tags linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
 * @frontendId 19
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head

	// 先将 p 节点移动 n
	p := dummy.Next
	for i := 0; i < n; i++ {
		p = p.Next
	}

	// 让节点 q 从头开始和 p 一起移动
	// p 移动到底时停止
	// p 最终走了 l - n 步, 也就是停在了倒数第 n 个节点
	prev := dummy
	q := dummy.Next
	for p != nil {
		p = p.Next
		prev = q
		q = q.Next
	}

	prev.Next = q.Next

	return dummy.Next
}
