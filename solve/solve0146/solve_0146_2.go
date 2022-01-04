package solve0146

type DoubleList2 struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleList2() *DoubleList2 {
	head, tail := NewNode(0, 0), NewNode(0, 0)
	head.Next = tail
	tail.Prev = head
	return &DoubleList2{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (dl *DoubleList2) Size() int {
	return dl.size
}

func (dl *DoubleList2) AddFirst(node *Node) {
	node.Prev = dl.head
	node.Next = dl.head.Next
	dl.head.Next.Prev = node
	dl.head.Next = node
	dl.size++
}

func (dl *DoubleList2) Remove(node *Node) {
	if node == nil || dl.size == 0 {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	dl.size--
}

func (dl *DoubleList2) RemoveLast() *Node {
	last := dl.tail.Prev
	dl.Remove(last)
	return last
}

func (dl *DoubleList2) MoveToHead(node *Node) {
	dl.Remove(node)
	dl.AddFirst(node)
}

type LRU2 struct {
	dl    *DoubleList2
	store map[int]*Node
	cap   int
}

func Constructor2(capacity int) *LRU2 {
	return &LRU2{
		cap:   capacity,
		dl:    NewDoubleList2(),
		store: make(map[int]*Node, capacity+1),
	}
}

func (lru *LRU2) Get(key int) int {
	node, ok := lru.store[key]
	if !ok {
		return -1
	}
	lru.dl.MoveToHead(node)
	return node.Val
}

func (lru *LRU2) Put(key, val int) {
	if node, ok := lru.store[key]; ok {
		node.Val = val
		lru.dl.MoveToHead(node)
	} else {
		n := NewNode(key, val)
		lru.dl.AddFirst(n)
		lru.store[key] = n

		if lru.dl.Size() > lru.cap {
			last := lru.dl.RemoveLast()
			delete(lru.store, last.Key)
		}
	}
}
