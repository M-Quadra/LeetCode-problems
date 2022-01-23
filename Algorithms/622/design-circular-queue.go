type MyCircularQueue struct {
	arr []int

	header int
	offset int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.arr[(this.header+this.offset)%len(this.arr)] = value
	this.offset++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.header++
	this.header %= len(this.arr)
	this.offset--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[this.header]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[(this.header+this.offset-1)%len(this.arr)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.offset <= 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.offset >= len(this.arr)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */