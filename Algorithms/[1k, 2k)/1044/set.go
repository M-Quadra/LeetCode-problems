func longestDupSubstring(s string) string {
	opt := ""
	for l, r := 1, len(s)+1; l < r; {
		set, m := map[string]struct{}{}, (l+r)/2
		for i := 0; i+m <= len(s); i++ {
			str := s[i : i+m]
			_, ok := set[str]
			if ok {
				opt = str
				break
			}

			set[str] = struct{}{}
		}

		if len(opt) >= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return opt
}