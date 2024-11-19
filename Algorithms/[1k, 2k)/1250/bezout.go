func isGoodArray(nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == 1
	}

	a := nums[0]
	for _, b := range nums[1:] {
		if a == 1 || b == 1 {
			return true
		}
		if a == b {
			continue
		}

		for a%b > 0 {
			c := a % b
			a = b
			b = c
		}
		a = b
	}
	return a == 1
}