type MyQueue struct {
	stackA []int
	stackB []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stackA = append(this.stackA, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackB) <= 0 {
		for len(this.stackA) > 0 {
			this.stackB = append(this.stackB, this.stackA[len(this.stackA)-1])
			this.stackA = this.stackA[:len(this.stackA)-1]
		}
	}

	opt := this.stackB[len(this.stackB)-1]
	this.stackB = this.stackB[:len(this.stackB)-1]
	return opt
}

func (this *MyQueue) Peek() int {
	if len(this.stackB) <= 0 {
		for len(this.stackA) > 0 {
			this.stackB = append(this.stackB, this.stackA[len(this.stackA)-1])
			this.stackA = this.stackA[:len(this.stackA)-1]
		}
	}
	return this.stackB[len(this.stackB)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackA) <= 0 && len(this.stackB) <= 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */