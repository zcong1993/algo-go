package solve0141

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 141
 * @title 环形链表
 * @difficulty 简单
 * @tags hash-table,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/linked-list-cycle/
 * @frontendId 141
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
