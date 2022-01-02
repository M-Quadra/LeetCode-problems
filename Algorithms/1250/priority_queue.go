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

func isGoodArray(nums []int) bool {
	pq, _ := New(&nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for !pq.Empty() {
		a := pq.Pop().(int)
		if a == 1 {
			return true
		}
		if pq.Empty() {
			break
		}

		b := a % pq.Top().(int)
		if b <= 0 {
			continue
		}
		if b == 1 {
			return true
		}
		pq.Push(b)
	}
	return false
}
