package solve0025

import (
	. "github.com/zcong1993/algo-go/pkg/linkedlist"
)

/**
 * @index 25
 * @title K 个一组翻转链表
 * @difficulty 困难
 * @tags recursion,linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
 * @frontendId 25
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverse 翻转 [head, end) 之间的链表.
func reverse(head, end *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != end {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 当前段为 [start, end)
	start, end := head, head
	// 元素少于 k 时, 返回原序
	for i := 0; i < k; i++ {
		if end == nil {
			return start
		}
		end = end.Next
	}
	// 翻转当前段
	newHead := reverse(start, end)
	// 之前段首 next 指向后续翻转后链表头
	start.Next = ReverseKGroup(end, k)

	return newHead
}
