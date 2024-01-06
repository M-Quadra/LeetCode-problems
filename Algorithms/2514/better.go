func countAnagrams(s string) int {
	const mod int = 1e9 + 7

	ans, tmp := 1, 1
	for _, str := range strings.Split(s, " ") {
		arr := [26]int{}
		for i, c := range str {
			i, c = i+1, c-'a'
			arr[c]++

			ans = (ans * i) % mod
			tmp = (tmp * arr[c]) % mod
		}
	}

	// 逆元
	inve := 1
	for a, b := tmp, mod-2; b != 0; b >>= 1 {
		if b&1 == 1 {
			inve = (a * inve) % mod
		}
		a = (a * a) % mod
	}
	ans = (ans * inve) % mod

	return ans
}
