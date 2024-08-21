func minFlipsMonoIncr(s string) int {
	opt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			opt++
		}
	}

	tmp := opt
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			tmp--
		} else {
			tmp++
		}

		if tmp < opt {
			opt = tmp
		}
	}
	return opt
}