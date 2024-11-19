func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ci, cj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if ci != cj {
			return ci < cj
		}
		return arr[i] < arr[j]
	})
	return arr
}