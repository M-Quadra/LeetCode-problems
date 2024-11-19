func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var _tree = make([]int, 4e5)

func findMax(i, l, r, wL, wR int) int {
	if wL <= l && r <= wR {
		return _tree[i]
	}

	m := (l + r) / 2
	if wR <= m {
		return findMax(i*2+1, l, m, wL, wR)
	} else if m < wL {
		return findMax(i*2+2, m+1, r, wL, wR)
	}
	return max(findMax(i*2+1, l, m, wL, m), findMax(i*2+2, m+1, r, m+1, wR))
}

func update(i, l, r, t, v int) {
	for {
		_tree[i] = max(_tree[i], v)
		if l == r {
			break
		}

		m := (l + r) / 2
		if t <= m {
			i = i*2 + 1
			r = m
		} else {
			i = i*2 + 2
			l = m + 1
		}
	}
}

func constrainedSubsetSum(nums []int, k int) int {
	for i := 0; i < len(nums)<<2; i++ {
		_tree[i] = -1e5
	}

	update(0, 0, len(nums)-1, 0, nums[0])
	for i := 1; i < len(nums); i++ {
		v := findMax(0, 0, len(nums)-1, max(0, i-k), i-1)
		v = max(v+nums[i], nums[i])
		update(0, 0, len(nums)-1, i, v)
	}
	return _tree[0]
}