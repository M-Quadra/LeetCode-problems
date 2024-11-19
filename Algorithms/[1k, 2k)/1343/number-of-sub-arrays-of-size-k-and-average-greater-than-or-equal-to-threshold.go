func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold *= k
	s := 0
	for i := 0; i < k; i++ {
		s += arr[i]
	}
	opt := 0
	if s >= threshold {
		opt++
	}

	for i := 0; k < len(arr); i, k = i+1, k+1 {
		s += arr[k] - arr[i]
		if s >= threshold {
			opt++
		}
	}
	return opt
}