func carPooling(trips [][]int, capacity int) bool {
	diffArr := [1e3 + 5]int{}
	for _, v := range trips {
		diffArr[v[1]] -= v[0]
		diffArr[v[2]] += v[0]
	}
	for _, v := range diffArr {
		capacity += v
		if capacity < 0 {
			return false
		}
	}
	return true
}