func groupThePeople(groupSizes []int) [][]int {
	sizeToGroup := map[int][]int{}
	res := [][]int{}
	for i, s := range groupSizes {
		arr := append(sizeToGroup[s], i)
		if len(arr) == s {
			res = append(res, arr)
			arr = nil
		}
		sizeToGroup[s] = arr
	}
	return res
}
