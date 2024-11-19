func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func suggestedProducts2(products []string, searchWord string) [][]string {
	opt := make([][]string, 0, len(searchWord))
	for i, searchChar := range searchWord {
		for j := 0; j < len(products); j++ {
			str := products[j]
			if i < len(str) && rune(str[i]) == searchChar {
				continue
			}

			products[j], products[len(products)-1] = products[len(products)-1], products[j]
			products = products[:len(products)-1]
			j--
		}

		sort.Strings(products)
		arr := make([]string, min(len(products), 3))
		copy(arr, products)
		opt = append(opt, arr)
	}
	return opt
}