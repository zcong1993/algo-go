package solve0908

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 908
 * @title 链表的中间结点
 * @difficulty 简单
 * @tags linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/middle-of-the-linked-list/
 * @frontendId 876
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针, 一个走两步, 一个走一步

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
