func longestSubsequence(arr []int, difference int) int {
	lenMap := map[int]int{}
	for _, v := range arr {
		lenMap[v] = lenMap[v-difference] + 1
	}

	opt := 0
	for _, v := range lenMap {
		if v > opt {
			opt = v
		}
	}
	return opt
}