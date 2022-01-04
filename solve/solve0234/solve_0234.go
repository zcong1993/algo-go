package solve0234

import (
	. "github.com/zcong1993/algo-go/pkg/linkedlist"
)

/**
 * @index 234
 * @title 回文链表
 * @difficulty 简单
 * @tags stack,recursion,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/palindrome-linked-list/
 * @frontendId 234
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	var arr []int
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}

	return true
}
