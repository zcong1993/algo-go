package list

type DoubleList2 struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleList2() *DoubleList2 {
	head := NewNode(nil)
	tail := NewNode(nil)
	head.Next = tail
	tail.Prev = head
	return &DoubleList2{head, tail, 0}
}

func (dl *DoubleList2) Head() *Node {
	if dl.size == 0 {
		return nil
	}
	return dl.head.Next
}

func (dl *DoubleList2) Tail() *Node {
	if dl.size == 0 {
		return nil
	}
	return dl.tail.Prev
}

func (dl *DoubleList2) AddFirst(node *Node) {
	node.Next = dl.head.Next
	node.Prev = dl.head
	dl.head.Next.Prev = node
	dl.head.Next = node
	dl.size++
}

func (dl *DoubleList2) AddLast(node *Node) {
	node.Next = dl.tail
	node.Prev = dl.tail.Prev
	dl.tail.Prev.Next = node
	dl.tail.Prev = node
	dl.size++
}

func (dl *DoubleList2) Remove(node *Node) {
	if node == nil || dl.size == 0 {
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	dl.size--
}

func (dl *DoubleList2) RemoveFirst() *Node {
	first := dl.Head()
	dl.Remove(first)
	return first
}

func (dl *DoubleList2) RemoveLast() *Node {
	last := dl.Tail()
	dl.Remove(last)
	return last
}

func (dl *DoubleList2) Size() int {
	return dl.size
}

func (dl *DoubleList2) Inspect() []interface{} {
	res := make([]interface{}, dl.size)
	cur := dl.Head()
	for i := 0; i < dl.size; i++ {
		res[i] = cur.Val
		cur = cur.Next
	}
	return res
}

var _ IDoubleList = &DoubleList2{}
