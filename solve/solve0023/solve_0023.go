package solve0023

import . "github.com/zcong1993/algo-go/pkg/linkedlist"

/**
 * @index 23
 * @title 合并K个升序链表
 * @difficulty 困难
 * @tags linked-list,divide-and-conquer,heap-priority-queue,merge-sort
 * @draft false
 * @link https://leetcode-cn.com/problems/merge-k-sorted-lists/
 * @frontendId 23
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type pq struct {
	store []*ListNode
	size  int
}

func newPQ() *pq {
	return &pq{
		store: []*ListNode{nil},
		size:  0,
	}
}

func parentIdx(idx int) int {
	return idx / 2
}

func leftIdx(idx int) int {
	return idx * 2
}

func rightIdx(idx int) int {
	return idx*2 + 1
}

func (p *pq) less(i, j int) bool {
	isLess := p.store[i].Val < p.store[j].Val
	return !isLess
}

func (p *pq) swap(i, j int) {
	p.store[i], p.store[j] = p.store[j], p.store[i]
}

func (p *pq) swim(idx int) {
	for idx > 1 {
		parent := parentIdx(idx)
		if p.less(idx, parent) {
			break
		}
		p.swap(idx, parent)
		idx = parent
	}
}

func (p *pq) sink(idx int) {
	for leftIdx(idx) <= p.size {
		bigIdx := leftIdx(idx)
		if right := rightIdx(idx); right <= p.size && p.less(bigIdx, right) {
			bigIdx = right
		}
		if p.less(bigIdx, idx) {
			break
		}
		p.swap(idx, bigIdx)
		idx = bigIdx
	}
}

func (p *pq) add(node *ListNode) {
	p.store = append(p.store, node)
	p.size++
	p.swim(p.size)
}

func (p *pq) removeTop() *ListNode {
	if p.size == 0 {
		return nil
	}

	top := p.store[1]
	p.swap(1, p.size)
	p.store = p.store[:p.size]
	p.size--
	p.sink(1)

	return top
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	p := newPQ()

	for _, h := range lists {
		cur := h
		for cur != nil {
			p.add(cur)
			cur = cur.Next
		}
	}

	dummy := &ListNode{Val: -1}
	cur := dummy
	top := p.removeTop()
	for top != nil {
		top.Next = nil
		cur.Next = top
		cur = cur.Next
		top = p.removeTop()
	}

	return dummy.Next
}
