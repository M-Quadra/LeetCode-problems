func makeSimilar(nums, target []int) int64 {
	cmp := func(a, b int) bool {
		iIsOdd, jIsOdd := a&1 == 1, b&1 == 1
		if iIsOdd == jIsOdd {
			return a < b
		}
		return iIsOdd
	}
	sort.Slice(nums, func(i, j int) bool {
		return cmp(nums[i], nums[j])
	})
	sort.Slice(target, func(i, j int) bool {
		return cmp(target[i], target[j])
	})

	var ans int64
	for i := range nums {
		ans += max(int64(nums[i]-target[i]), 0) / 2
	}
	return ans
}
