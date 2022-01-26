func finalPrices(prices []int) []int {
	opt := make([]int, len(prices))
	for i := range prices {
		v := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] && prices[j] > v {
				v = prices[j]
				break
			}
		}
		opt[i] = prices[i] - v
	}
	return opt
}