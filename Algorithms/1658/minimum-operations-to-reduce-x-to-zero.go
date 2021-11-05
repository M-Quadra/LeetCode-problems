func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minOperations(nums []int, x int) int {
	l, s := -1, 0
	for i, v := range nums {
		s += v
		if s >= x {
			l = i
			break
		}
	}
	if s < x {
		return -1
	}

	opt := 0x7FFFFF
	if s == x {
		opt = l + 1
	}

	r := len(nums)
	for ; l >= 0; l-- {
		for s -= nums[l]; l <= r-1 && s < x; r-- {
			s += nums[r-1]
		}
		if s == x {
			opt = min(opt, l+len(nums)-r)
		}
	}

	if opt == 0x7FFFFF {
		opt = -1
	}
	return opt
}