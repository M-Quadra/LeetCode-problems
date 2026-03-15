func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0], res[len(res)-1] = 1, 1
	for i := 1; i <= rowIndex/2; i++ {
		res[i] = res[i-1] * (len(res) - i) / i
		res[rowIndex-i] = res[i]
	}
	return res
}
