package queue

import (
	"errors"

	"github.com/zcong1993/algo-go/pkg/list"
)

var ErrNotExists = errors.New("not exists")

type IDeQueue interface {
	AddFirst(data interface{})
	AddLast(data interface{})
	PollFirst() (interface{}, error)
	PollLast() (interface{}, error)
	PeekFirst() (interface{}, error)
	PeekLast() (interface{}, error)
	Size() interface{}
}

type DeQueue struct {
	ll *list.DoubleList
}

func NewDeQueue() *DeQueue {
	return &DeQueue{ll: list.NewDoubleList()}
}

func (q *DeQueue) AddFirst(data interface{}) {
	q.ll.AddFirst(&list.Node{Val: data})
}

func (q *DeQueue) AddLast(data interface{}) {
	q.ll.AddLast(&list.Node{Val: data})
}

func (q *DeQueue) PollFirst() (interface{}, error) {
	first := q.ll.Head()
	if first == nil {
		return -1, ErrNotExists
	}
	q.ll.RemoveFirst()
	return first.Val, nil
}

func (q *DeQueue) PollLast() (interface{}, error) {
	last := q.ll.Tail()
	if last == nil {
		return -1, ErrNotExists
	}
	q.ll.RemoveLast()
	return last.Val, nil
}

func (q *DeQueue) PeekFirst() (interface{}, error) {
	first := q.ll.Head()
	if first == nil {
		return -1, ErrNotExists
	}
	return first.Val, nil
}

func (q *DeQueue) PeekLast() (interface{}, error) {
	last := q.ll.Tail()
	if last == nil {
		return -1, ErrNotExists
	}
	return last.Val, nil
}

func (q *DeQueue) Size() interface{} {
	return q.ll.Size()
}

var _ IDeQueue = &DeQueue{}
