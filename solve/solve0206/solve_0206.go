package solve0206

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 206
 * @title 反转链表
 * @difficulty 简单
 * @tags recursion,linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/reverse-linked-list/
 * @frontendId 206
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	node := head

	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	last := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
