package solve0160

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 160
 * @title 相交链表
 * @difficulty 简单
 * @tags hash-table,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
 * @frontendId 160
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	curA := headA
	for curA != nil {
		mp[curA] = struct{}{}
		curA = curA.Next
	}

	curB := headB
	for curB != nil {
		if _, ok := mp[curB]; ok {
			return curB
		}
		curB = curB.Next
	}

	return nil
}

func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}
