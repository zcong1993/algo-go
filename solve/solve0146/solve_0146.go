package solve0146

/**
 * @index 146
 * @title LRU 缓存机制
 * @difficulty 中等
 * @tags design,hash-table,linked-list,doubly-linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/lru-cache/
 * @frontendId 146
 */

type Node struct {
	Prev *Node
	Next *Node
	Val  int
	Key  int
}

func NewNode(key, val int) *Node {
	return &Node{Val: val, Key: key}
}

type IDoublelist interface {
	AddFirst(node *Node)
	Remove(node *Node)
	RemoveLast() *Node
	MoveToHead(node *Node)
	Size() int
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleList() *DoubleList {
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.Next = tail
	tail.Prev = head
	return &DoubleList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (dl *DoubleList) Size() int {
	return dl.size
}

func (dl *DoubleList) AddFirst(node *Node) {
	node.Next = dl.head.Next
	node.Prev = dl.head
	dl.head.Next.Prev = node
	dl.head.Next = node
	dl.size++
}

func (dl *DoubleList) Remove(node *Node) {
	if node == nil || dl.size == 0 {
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	dl.size--
}

func (dl *DoubleList) RemoveLast() *Node {
	last := dl.tail.Prev
	dl.Remove(last)
	return last
}

func (dl *DoubleList) MoveToHead(node *Node) {
	dl.Remove(node)
	dl.AddFirst(node)
}

type ILRU interface {
	Get(key int) int
	Put(key, val int)
}

type LRUCache struct {
	dl       *DoubleList
	store    map[int]*Node
	capacity int
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		dl:       NewDoubleList(),
		store:    make(map[int]*Node, capacity+1),
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.store[key]
	if !ok {
		return -1
	}
	this.dl.MoveToHead(n)
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.store[key]; ok {
		n.Val = value
		this.dl.MoveToHead(n)
	} else {
		node := &Node{
			Val: value,
			Key: key,
		}
		this.dl.AddFirst(node)
		this.store[key] = node
		if this.dl.Size() > this.capacity {
			last := this.dl.RemoveLast()
			delete(this.store, last.Key)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
