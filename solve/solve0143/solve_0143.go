package solve0143

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 143
 * @title 重排链表
 * @difficulty 中等
 * @tags stack,recursion,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/reorder-list/
 * @frontendId 143
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. 找到中点
	fast, slow := head, head
	var leftEnd *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		leftEnd = slow
		slow = slow.Next
	}
	// 如果是奇数, 需要在向前移动一步
	// 因为左边需要比右边长
	if fast != nil {
		leftEnd = slow
		slow = slow.Next
	}

	// 2. 切断为左右两部分
	right := slow
	leftEnd.Next = nil

	// 3. 翻转右半边
	right = reverse(right)

	// 4. 将右边链表节点按顺序插入左边节点之间
	curL, curR := head, right
	for curL != nil && curR != nil {
		nextL, nextR := curL.Next, curR.Next
		curR.Next = nextL
		curL.Next = curR
		curL, curR = nextL, nextR
	}
}
