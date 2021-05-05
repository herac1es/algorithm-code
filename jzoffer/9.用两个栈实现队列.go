package jzoffer

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

func ConstructorCQ() CQueue {
	return CQueue{
		stack1: make(Stack, 0, 10),
		stack2: make(Stack, 0, 10),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Size() == 0 && this.stack1.Size() == 0 {
		return -1
	}
	if this.stack2.Size() == 0 {
		for this.stack1.Size() > 0 {
			this.stack2.Push(this.stack1.Pop())
		}
	}
	return this.stack2.Pop()
}

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if s.Size() == 0 {
		return -1
	}
	r := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return r
}

func (s *Stack) Size() int {
	return len(*s)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
