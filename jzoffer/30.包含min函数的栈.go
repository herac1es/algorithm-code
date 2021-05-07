package jzoffer

type MinStack struct {
	buf [][2]int // [0]: 元素, [1]:当前元素入栈时的最小值
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		buf: make([][2]int, 0, 512),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.buf) == 0 {
		this.buf = append(this.buf, [2]int{x, x})
	} else {
		min := this.buf[len(this.buf)-1][1]
		if x < min {
			min = x
		}
		this.buf = append(this.buf, [2]int{x, min})
	}
}

func (this *MinStack) Pop() {
	if len(this.buf) == 0 {
		return
	}
	this.buf = this.buf[:len(this.buf)-1]
}

func (this *MinStack) Top() int {
	if len(this.buf) == 0 {
		return -1
	}
	return this.buf[len(this.buf)-1][0]
}

func (this *MinStack) Min() int {
	if len(this.buf) == 0 {
		return -1
	}
	return this.buf[len(this.buf)-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
