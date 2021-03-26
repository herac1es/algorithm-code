package classic

type PriorityQueue struct {
	arr  []int
	size int
}

func New() PriorityQueue {
	return PriorityQueue{
		arr:  make([]int, 1),
		size: 0,
	}
}

// 上浮
func (pq *PriorityQueue) swim(idx int) {
	for idx > 1 && pq.less(idx, idx>>1) {
		pq.arr[idx], pq.arr[idx>>1] = pq.arr[idx>>1], pq.arr[idx]
		idx >>= 1
	}
}

func (pq *PriorityQueue) less(i, j int) bool {
	return pq.arr[i] > pq.arr[j]
}

// 下沉
func (pq *PriorityQueue) sink(idx int) {
	for (idx << 1) <= pq.size {
		t := idx << 1
		if t+1 <= pq.size && pq.less(t+1, t) {
			t = t + 1
		}
		if !pq.less(t, idx) {
			break
		}
		pq.arr[t], pq.arr[idx] = pq.arr[idx], pq.arr[t]
		idx = t
	}
}

func (pq *PriorityQueue) Pop() int {
	if pq.size == 0 {
		return 0
	}
	ret := pq.arr[1]
	pq.arr[1] = pq.arr[pq.size]
	pq.arr = pq.arr[:pq.size]
	pq.size--
	pq.sink(1)
	return ret
}

func (pq *PriorityQueue) Push(item int) {
	pq.size++
	pq.arr = append(pq.arr, item)
	idx := pq.size
	pq.swim(idx)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue) List() []int {
	ret := make([]int, 0, len(pq.arr))
	ret = append(ret, pq.arr...)
	return ret[1:]
}
