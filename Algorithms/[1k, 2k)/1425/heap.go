func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var _heap []int

func constrainedSubsetSum(nums []int, k int) int {
	n := int(math.Ceil(math.Log2(float64(k))) + 0.5)
	s := (1 << (n + 1)) - 1
	if len(_heap) < s {
		_heap = make([]int, s)
	}
	for i := 0; i < s; i++ {
		_heap[i] = -1e5
	}
	s -= 1 << n

	opt := int(-1e5)
	for i, v := range nums {
		v = max(_heap[0]+v, v)
		opt = max(opt, v)

		i = i%k + s
		_heap[i] = v
		for i = (i - 1) >> 1; i >= 0; i = (i - 1) >> 1 {
			v = max(_heap[i<<1|1], _heap[(i+1)<<1])
			if v == _heap[i] {
				break
			}
			_heap[i] = v
		}
	}

	return opt
}