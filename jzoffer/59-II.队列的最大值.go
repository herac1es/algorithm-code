package jzoffer

type MaxQueue struct {
	maxStack []int
	valQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		maxStack: make([]int, 0),
		valQueue: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.valQueue) == 0 || len(this.maxStack) == 0 {
		return -1
	}
	return this.maxStack[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.valQueue = append(this.valQueue, value)
	for len(this.maxStack) > 0 && this.maxStack[len(this.maxStack)-1] < value {
		this.maxStack = this.maxStack[:len(this.maxStack)-1]
	}
	this.maxStack = append(this.maxStack, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.valQueue) == 0 {
		return -1
	}
	value := this.valQueue[0]
	this.valQueue = this.valQueue[1:]
	if this.maxStack[0] == value {
		this.maxStack = this.maxStack[1:]
	}
	return value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
