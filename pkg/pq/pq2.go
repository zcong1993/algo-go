package pq

type PQ2 struct {
	store []int
	cap   int
	size  int
	t     T
}

func NewPQ2(cap int, t T) *PQ2 {
	return &PQ2{
		store: make([]int, cap+1),
		cap:   cap,
		size:  0,
		t:     t,
	}
}

func (pq *PQ2) parent(idx int) int {
	return idx / 2
}

func (pq *PQ2) left(idx int) int {
	return idx * 2
}

func (pq *PQ2) right(idx int) int {
	return idx*2 + 1
}

func (pq *PQ2) less(i, j int) bool {
	isLess := pq.store[i] < pq.store[j]
	if pq.t == MAX {
		return isLess
	} else {
		return !isLess
	}
}

func (pq *PQ2) swap(i, j int) {
	pq.store[i], pq.store[j] = pq.store[j], pq.store[i]
}

// 如果某个节点 A 比它的父节点大，那么 A 不应该做子节点，应该把父节点换下来，自己去做父节点，这就是对 A 的上浮.
func (pq *PQ2) swim(idx int) {
	for idx > 1 {
		parent := pq.parent(idx)
		// parent < current
		if pq.less(parent, idx) {
			pq.swap(parent, idx)
			idx = parent
		} else {
			break
		}
	}
}

// 如果某个节点 A 比它的子节点（中的一个）小，那么 A 就不配做父节点，应该下去，下面那个更大的节点上来做父节点，这就是对 A 进行下沉.
func (pq *PQ2) sink(idx int) {
	for pq.left(idx) <= pq.size {
		bigIdx := pq.left(idx)
		if right := pq.right(idx); right <= pq.size && pq.less(bigIdx, right) {
			bigIdx = right
		}
		if pq.less(bigIdx, idx) {
			break
		}
		pq.swap(bigIdx, idx)
		idx = bigIdx
	}
}

func (pq *PQ2) Size() int {
	return pq.size
}

func (pq *PQ2) Add(val int) {
	pq.size++
	pq.store[pq.size] = val
	pq.swim(pq.size)
}

func (pq *PQ2) DelTop() int {
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

func (pq *PQ2) Peek() int {
	if pq.size == 0 {
		panic("nil pq")
	}
	return pq.store[1]
}

var _ IPQ = &PQ2{}
