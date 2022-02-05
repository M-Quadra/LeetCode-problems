func singleNumber(nums []int) []int {
	c := 0
	for _, v := range nums {
		c ^= v
	}
	lowBit := c & -c

	a, b := 0, 0
	for _, v := range nums {
		if v&lowBit <= 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}