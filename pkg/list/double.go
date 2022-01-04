package list

type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

type IDoubleList interface {
	Head() *Node
	Tail() *Node
	AddFirst(node *Node)
	AddLast(node *Node)
	RemoveFirst() *Node
	RemoveLast() *Node
	Remove(node *Node)
	Size() int
}

type DoubleList struct {
	head *Node
	tail *Node
	s    int
}

func NewDoubleList() *DoubleList {
	head := &Node{Val: 0}
	tail := &Node{Val: 0}
	head.Next = tail
	tail.Prev = head
	return &DoubleList{
		head: head,
		tail: tail,
		s:    0,
	}
}

func (dl *DoubleList) Head() *Node {
	if dl.s == 0 {
		return nil
	}
	return dl.head.Next
}

func (dl *DoubleList) Tail() *Node {
	if dl.s == 0 {
		return nil
	}
	return dl.tail.Prev
}

func (dl *DoubleList) AddFirst(node *Node) {
	node.Next = dl.head.Next
	node.Prev = dl.head
	dl.head.Next.Prev = node
	dl.head.Next = node
	dl.s++
}

func (dl *DoubleList) AddLast(node *Node) {
	node.Prev = dl.tail.Prev
	node.Next = dl.tail
	dl.tail.Prev.Next = node
	dl.tail.Prev = node
	dl.s++
}

func (dl *DoubleList) RemoveFirst() *Node {
	head := dl.Head()
	dl.Remove(head)
	return head
}

func (dl *DoubleList) RemoveLast() *Node {
	tail := dl.Tail()
	dl.Remove(tail)
	return tail
}

func (dl *DoubleList) Remove(node *Node) {
	if node == nil || dl.s == 0 {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	dl.s--
}

func (dl *DoubleList) Size() int {
	return dl.s
}

func (dl *DoubleList) Inspect() []interface{} {
	res := make([]interface{}, dl.s)
	i := 0
	cur := dl.Head()
	for i < dl.s {
		res[i] = cur.Val
		cur = cur.Next
		i++
	}
	return res
}

var _ IDoubleList = &DoubleList{}
