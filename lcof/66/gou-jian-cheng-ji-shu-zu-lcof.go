func constructArr(a []int) []int {
	if len(a) <= 0 {
		return a
	}

	b := make([]int, len(a))
	b[len(b)-1] = a[len(a)-1]
	for i := len(b) - 2; i >= 0; i-- {
		b[i] = a[i] * b[i+1]
	}
	for i := 1; i < len(a); i++ {
		a[i] *= a[i-1]
	}

	for i := range b {
		l := 1
		if i-1 >= 0 {
			l = a[i-1]
		}
		r := 1
		if i+1 < len(b) {
			r = b[i+1]
		}
		b[i] = l * r
	}

	return b
}
