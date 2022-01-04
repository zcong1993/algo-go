package pq

type T = int

const (
	MAX T = iota + 1
	MIN
)

type IPQ interface {
	Size() int
	Add(val int)
	DelTop() int
	Peek() int
}

type PQ struct {
	store []int
	size  int
	cap   int
	t     T
}

func NewPQ(cap int, t T) *PQ {
	return &PQ{
		store: make([]int, cap+1),
		size:  0,
		cap:   cap,
		t:     t,
	}
}

func parent(idx int) int {
	return idx / 2
}

func left(idx int) int {
	return idx * 2
}

func right(idx int) int {
	return idx*2 + 1
}

func (pq *PQ) less(i, j int) bool {
	isLess := pq.store[i] < pq.store[j]
	if pq.t == MAX {
		return isLess
	} else {
		return !isLess
	}
}

func (pq *PQ) swap(i, j int) {
	pq.store[i], pq.store[j] = pq.store[j], pq.store[i]
}

func (pq *PQ) swim(idx int) {
	for idx > 1 {
		parentIdx := parent(idx)
		if pq.less(idx, parentIdx) {
			break
		}
		pq.swap(parentIdx, idx)
		idx = parentIdx
	}
}

func (pq *PQ) sink(idx int) {
	for left(idx) <= pq.size {
		bigIdx := left(idx)
		if rightIdx := right(idx); rightIdx <= pq.size && pq.less(bigIdx, rightIdx) {
			bigIdx = rightIdx
		}
		if pq.less(bigIdx, idx) {
			break
		}
		pq.swap(bigIdx, idx)
		idx = bigIdx
	}
}

func (pq *PQ) Size() int {
	return pq.size
}

func (pq *PQ) Add(val int) {
	pq.size++
	pq.store[pq.size] = val
	pq.swim(pq.size)
}

func (pq *PQ) DelTop() int {
	if pq.size == 0 {
		panic("nil pq")
	}
	top := pq.store[1]
	pq.swap(1, pq.size)
	pq.store[pq.size] = -1
	pq.size--
	pq.sink(1)
	return top
}

func (pq *PQ) Peek() int {
	if pq.size == 0 {
		panic("nil pq")
	}
	return pq.store[1]
}

var _ IPQ = &PQ{}
