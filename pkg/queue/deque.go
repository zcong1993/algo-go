package queue

import (
	"errors"

	"github.com/zcong1993/algo-go/pkg/list"
)

var ErrNotExists = errors.New("not exists")

type IDeQueue interface {
	AddFirst(data int)
	AddLast(data int)
	PollFirst() (int, error)
	PollLast() (int, error)
	PeekFirst() (int, error)
	PeekLast() (int, error)
	Size() int
}

type DeQueue struct {
	ll *list.DoubleList
}

func NewDeQueue() *DeQueue {
	return &DeQueue{ll: list.NewDoubleList()}
}

func (q *DeQueue) AddFirst(data int) {
	q.ll.AddFirst(&list.Node{Val: data})
}

func (q *DeQueue) AddLast(data int) {
	q.ll.AddLast(&list.Node{Val: data})
}

func (q *DeQueue) PollFirst() (int, error) {
	first := q.ll.Head()
	if first == nil {
		return -1, ErrNotExists
	}
	q.ll.RemoveFirst()
	return first.Val, nil
}

func (q *DeQueue) PollLast() (int, error) {
	last := q.ll.Tail()
	if last == nil {
		return -1, ErrNotExists
	}
	q.ll.RemoveLast()
	return last.Val, nil
}

func (q *DeQueue) PeekFirst() (int, error) {
	first := q.ll.Head()
	if first == nil {
		return -1, ErrNotExists
	}
	return first.Val, nil
}

func (q *DeQueue) PeekLast() (int, error) {
	last := q.ll.Tail()
	if last == nil {
		return -1, ErrNotExists
	}
	return last.Val, nil
}

func (q *DeQueue) Size() int {
	return q.ll.Size()
}
