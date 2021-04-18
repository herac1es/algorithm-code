package leetcode

// 155.最小栈
type MinStack struct {
	c    [][2]int // 数组0元素为当前元素，1元素为当前元素为栈顶时的最小值
	size int
}

/** initialize your data structure here. */
func Constructor_() MinStack {
	return MinStack{
		c:    make([][2]int, 0),
		size: 0,
	}
}

func (this *MinStack) Push(val int) {
	if this.size == 0 {
		this.c = append(this.c, [2]int{val, val})
	} else {
		node := [2]int{val, val}
		if v := this.c[this.size-1][1]; v < val {
			node[1] = v
		}
		this.c = append(this.c, node)
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.c = this.c[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.c[this.size-1][0]
}

func (this *MinStack) GetMin() int {
	return this.c[this.size-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
