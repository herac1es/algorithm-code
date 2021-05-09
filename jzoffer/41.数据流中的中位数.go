package jzoffer

type MedianFinder struct {
	maxer Heap // a min top heap
	miner Heap // a max top heap
}

type Heap struct {
	buf  []int
	size int
	less func(i, j int) bool
}

// default: a max top heap
func NewHeap(less func(i, j int) bool) Heap {
	return Heap{
		buf:  make([]int, 1),
		size: 0,
		less: less,
	}
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Push(val int) {
	h.buf = append(h.buf, val)
	h.size++
	h.swim(h.size)
}

func (h *Heap) Pop() int {
	if h.size <= 0 {
		return -1
	}
	ret := h.buf[1]
	h.buf[1], h.buf[h.size] = h.buf[h.size], h.buf[1]
	h.buf = h.buf[:h.size]
	h.size--
	h.sink(1)
	return ret
}

func (h *Heap) Top() int {
	if h.Size() <= 0 {
		return -1
	}
	return h.buf[1]
}

func (h *Heap) swim(idx int) {
	for idx > 1 && h.less(h.buf[idx/2], h.buf[idx]) {
		h.buf[idx/2], h.buf[idx] = h.buf[idx], h.buf[idx/2]
		idx /= 2
	}
}

func (h *Heap) sink(idx int) {
	for idx*2 <= h.size {
		j := idx * 2
		if j+1 <= h.size && h.less(h.buf[j], h.buf[j+1]) {
			j = j + 1
		}
		// if buf[j] <= buf[idx]
		if !h.less(h.buf[idx], h.buf[j]) {
			return
		}
		h.buf[idx], h.buf[j] = h.buf[j], h.buf[idx]
		idx = j
	}
}

/** initialize your data structure here. */
func Constructor41() MedianFinder {
	return MedianFinder{
		maxer: NewHeap(func(i, j int) bool {
			return j < i
		}),
		miner: NewHeap(func(i, j int) bool {
			return i < j
		}),
	}
}

// 时间: O(logn)
func (this *MedianFinder) AddNum(num int) {
	if this.maxer.Size() == this.miner.Size() {
		this.miner.Push(num)
		this.maxer.Push(this.miner.Pop())
	} else {
		this.maxer.Push(num)
		this.miner.Push(this.maxer.Pop())
	}
}

// 时间: O(1)
func (this *MedianFinder) FindMedian() float64 {
	if this.maxer.Size() == this.miner.Size() {
		return float64(this.maxer.Top()+this.miner.Top()) / 2
	} else {
		return float64(this.maxer.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
