func maxPalindromes(s string, k int) int {
	isPalindrome := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	arr := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		for j := i + k - 1; j < len(s); j++ {
			if !isPalindrome(s[i : j+1]) {
				continue
			}
			arr[j+1] = max(arr[j+1], arr[i]+1)
			break
		}
		arr[i+1] = max(arr[i+1], arr[i])
	}
	return arr[len(arr)-1]
}
