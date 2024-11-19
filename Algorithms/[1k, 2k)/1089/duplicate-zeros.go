func duplicateZeros(arr []int) {
	c := 0
	for _, v := range arr {
		if v == 0 {
			c++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			c--
		}
		j := i + c
		if j >= len(arr) {
			continue
		}

		arr[j] = arr[i]
		if arr[i] == 0 && j+1 < len(arr) {
			arr[j+1] = 0
		}
	}
}