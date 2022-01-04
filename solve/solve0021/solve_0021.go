package solve0021

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 21
 * @title 合并两个有序链表
 * @difficulty 简单
 * @tags recursion,linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/merge-two-sorted-lists/
 * @frontendId 21
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1. 和合并有序数组一样
// 2. 每个链表元素只会遍历一次, 可以直接复用

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curD := dummy
	cur1, cur2 := list1, list2
	// 两边都有, 每次前进小的那条
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			curD.Next = cur1
			cur1 = cur1.Next
		} else {
			curD.Next = cur2
			cur2 = cur2.Next
		}
		curD = curD.Next
	}

	// 有剩余的直接拼在后面
	for cur1 != nil {
		curD.Next = cur1
		cur1 = cur1.Next
		curD = curD.Next
	}

	for cur2 != nil {
		curD.Next = cur2
		cur2 = cur2.Next
		curD = curD.Next
	}

	return dummy.Next
}
