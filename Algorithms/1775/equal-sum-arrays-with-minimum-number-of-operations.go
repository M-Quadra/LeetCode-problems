func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOperations(nums1 []int, nums2 []int) int {
	diff, cntArr := 0, make([]int, 6)
	for _, v := range nums1 {
		diff -= v
		cntArr[v-1]++
	}
	for _, v := range nums2 {
		diff += v
		cntArr[6-v]++
	}
	if diff < 0 {
		diff *= -1
		for i, j := 0, len(cntArr)-1; i < j; i, j = i+1, j-1 {
			cntArr[i], cntArr[j] = cntArr[j], cntArr[i]
		}
	}

	opt := 0
	for i := 0; i < 5 && diff > 0; i++ {
		c := min(cntArr[i], diff/(5-i))
		opt += c

		diff -= c * (5 - i)
		cntArr[i+1] += cntArr[i] - c
	}

	if diff > 0 {
		return -1
	}
	return opt
}