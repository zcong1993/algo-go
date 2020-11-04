package list

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type IDoubleList interface {
	Head() *Node
	Tail() *Node
	AddFirst(node *Node)
	AddLast(node *Node)
	RemoveFirst()
	RemoveLast()
	Remove(node *Node)
	Size() int
}

type DoubleList struct {
	head *Node
	tail *Node
	s    int
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		head: nil,
		tail: nil,
		s:    0,
	}
}

func (dl *DoubleList) Head() *Node {
	return dl.head
}

func (dl *DoubleList) Tail() *Node {
	return dl.tail
}

func (dl *DoubleList) AddFirst(node *Node) {
	if dl.s == 0 {
		dl.head = node
		dl.tail = node
		dl.s = 1
		return
	}
	dl.head.Prev = node
	node.Next = dl.head
	dl.head = node
	dl.s++
}

func (dl *DoubleList) AddLast(node *Node) {
	if dl.s == 0 {
		dl.head = node
		dl.tail = node
		dl.s = 1
		return
	}
	dl.tail.Next = node
	node.Prev = dl.tail
	dl.tail = node
	dl.s++
}

func (dl *DoubleList) RemoveFirst() {
	if dl.s == 0 {
		return
	}
	if dl.s == 1 {
		dl.head = nil
		dl.tail = nil
		dl.s = 0
		return
	}
	if dl.head.Next != nil {
		dl.head.Next.Prev = nil
	}
	dl.head = dl.head.Next
	dl.s--
}

func (dl *DoubleList) RemoveLast() {
	if dl.s == 0 {
		return
	}
	if dl.s == 1 {
		dl.head = nil
		dl.tail = nil
		dl.s = 0
		return
	}
	if dl.tail.Prev != nil {
		dl.tail.Prev.Next = nil
	}
	dl.tail = dl.tail.Prev
	dl.s--
}

func (dl *DoubleList) Remove(node *Node) {
	if node == nil || dl.s == 0 {
		return
	}
	if dl.s == 1 && node == dl.head {
		dl.head = nil
		dl.tail = nil
		dl.s = 0
		return
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	dl.s--
}

func (dl *DoubleList) Size() int {
	return dl.s
}

func (dl *DoubleList) Inspect() []int {
	res := make([]int, dl.s)
	i := 0
	cur := dl.head
	for cur != nil {
		res[i] = cur.Val
		cur = cur.Next
		i++
	}
	return res
}
