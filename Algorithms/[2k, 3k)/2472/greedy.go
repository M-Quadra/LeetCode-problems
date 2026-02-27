func maxPalindromes(s string, k int) int {
	isPalindrome := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	ans := 0
	for i := 0; i <= len(s)-k; {
		if j := i + k; isPalindrome(s[i:j]) {
			i = j
		} else if j := i + k + 1; j <= len(s) && isPalindrome(s[i:j]) {
			i = j
		} else {
			i++
			continue
		}
		ans++
	}
	return ans
}
