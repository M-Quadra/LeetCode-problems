func arithmeticTriplets(nums []int, diff int) int {
	arr := make([]bool, 201)
	for _, v := range nums {
		arr[v] = true
	}

	ans := 0
	for i := 1; i < 201-diff; i++ {
		if arr[i] && arr[i+diff] && arr[i+diff+diff] {
			ans++
		}
	}
	return ans
}
