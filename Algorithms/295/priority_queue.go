// PriorityQueue std::priority_queue
type PriorityQueue struct {
	sub subHeap
}

// Size len(slice)
func (pq PriorityQueue) Size() int {
	return pq.sub.Len()
}

// Empty len(slice) == 0
func (pq PriorityQueue) Empty() bool {
	return pq.Size() <= 0
}

// Push priority_queue.push(x)
func (pq PriorityQueue) Push(x interface{}) {
	heap.Push(pq.sub, x)
}

// Pop priority_queue.pop() with result
func (pq PriorityQueue) Pop() interface{} {
	if pq.sub.Len() <= 0 {
		return nil
	}

	return heap.Pop(pq.sub)
}

// Top slice[0]
func (pq PriorityQueue) Top() interface{} {
	if pq.sub.Len() <= 0 {
		return nil
	}

	return pq.sub.slicePtr.Elem().Index(0).Interface()
}

// New need point of sliceP
//  less(i, j int) { return a[i]<b[i] } is min-heap
func New(slicePtr interface{}, less func(i, j int) bool) (*PriorityQueue, error) {
	rv := reflect.ValueOf(slicePtr)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		return nil, errors.New("slicePtr must be a slice pointer")
	}
	if less == nil {
		return nil, errors.New("less must not be nil")
	}

	opt := &PriorityQueue{
		sub: subHeap{
			slicePtr: rv,
			less:     less,
		},
	}
	heap.Init(opt.sub)
	return opt, nil
}

type subHeap struct {
	heap.Interface

	slicePtr reflect.Value
	less     func(i, j int) bool
}

func (sh subHeap) Len() int {
	return sh.slicePtr.Elem().Len()
}

func (sh subHeap) Swap(i, j int) {
	reflect.Swapper(sh.slicePtr.Elem().Interface())(i, j)
}

func (sh subHeap) Less(i, j int) bool {
	return sh.less(i, j)
}

func (sh subHeap) Push(x interface{}) {
	slice := reflect.Append(sh.slicePtr.Elem(), reflect.ValueOf(x))
	sh.slicePtr.Elem().Set(slice)
}

func (sh subHeap) Pop() interface{} {
	opt := sh.slicePtr.Elem().Index(sh.Len() - 1).Interface()
	sh.slicePtr.Elem().SetLen(sh.Len() - 1)
	return opt
}

type MedianFinder struct {
	lPq, rPq *PriorityQueue
}

func Constructor() MedianFinder {
	lArr := []int{}
	rArr := []int{}

	lPq, _ := New(&lArr, func(i, j int) bool {
		return lArr[i] > lArr[j]
	})
	rPq, _ := New(&rArr, func(i, j int) bool {
		return rArr[i] < rArr[j]
	})
	return MedianFinder{
		lPq: lPq,
		rPq: rPq,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.lPq.Size() <= this.rPq.Size() {
		this.lPq.Push(num)
	} else {
		this.rPq.Push(num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	for !this.lPq.Empty() && !this.rPq.Empty() && this.lPq.Top().(int) > this.rPq.Top().(int) {
		this.lPq.Push(this.rPq.Pop())
		this.rPq.Push(this.lPq.Pop())
	}

	if this.lPq.Size() != this.rPq.Size() {
		return float64(this.lPq.Top().(int))
	}

	return float64(this.lPq.Top().(int))/2 + float64(this.rPq.Top().(int))/2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
