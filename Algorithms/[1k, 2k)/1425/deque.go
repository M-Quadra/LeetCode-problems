func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func constrainedSubsetSum(nums []int, k int) int {
	deque := list.New()
	deque.PushBack([2]int{-1, -1e5})

	front := func() [2]int {
		return deque.Front().Value.([2]int)
	}
	back := func() [2]int {
		return deque.Back().Value.([2]int)
	}

	opt := int(-1e5)
	for i, v := range nums {
		for front()[0]+k < i {
			deque.Remove(deque.Front())
		}

		v = max(front()[1]+v, v)
		opt = max(opt, v)

		for deque.Len() > 0 && back()[1] <= v {
			deque.Remove(deque.Back())
		}
		deque.PushBack([2]int{i, v})
	}
	return opt
}