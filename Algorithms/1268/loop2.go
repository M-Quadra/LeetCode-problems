func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	to := make([]int, len(products))
	for i := range to {
		to[i] = i + 1
	}
	to[len(to)-1] = -1

	opt := make([][]string, len(searchWord))
	st, lenPro := 0, len(products)
	for i, searchChar := range searchWord {
		for j, last := st, st; j != -1; j = to[j] {
			str := products[j]
			if i < len(str) && rune(str[i]) == searchChar {
				opt[i] = append(opt[i], str)
				last = j
				continue
			}

			lenPro--
			to[last] = to[j]
			if j == st {
				st = to[j]
			}
		}

		arr := make([]string, 0, min(lenPro, 3))
		for j := st; j != -1 && len(arr) < 3; j = to[j] {
			arr = append(arr, products[j])
		}
		opt[i] = arr
	}
	return opt
}