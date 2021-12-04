func log2(a int) int {
	return bits.UintSize - bits.LeadingZeros(uint(a)) - 1
}

func findKthBit(n, k int) byte {
	opt := byte(0)
	for n = log2(k); k > 1; n = log2(k) {
		if k == 1<<n {
			opt ^= 1
			break
		}

		n++
		k = (1 << n) - k
		opt ^= 1
	}
	return '0' + opt
}