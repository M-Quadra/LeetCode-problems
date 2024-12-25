func taskSchedulerII(tasks []int, space int) int64 {
	dayMap := map[int]int64{}
	ans := int64(0)
	for _, v := range tasks {
		ans = max(ans, dayMap[v]) + 1
		dayMap[v] = ans + int64(space)
	}
	return ans
}
