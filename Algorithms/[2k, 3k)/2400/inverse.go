const _mod = 1_000_000_000 + 7

var _factorials = make([]int, 1_001)

func init() {
	_factorials[0] = 1
	for i := 1; i <= 1_000; i++ {
		_factorials[i] = i * _factorials[i-1] % _mod
	}
}

func binpow(a, b int) (res int) {
	for res = 1; b > 0; a, b = (a*a)%_mod, b>>1 {
		if b&1 == 1 {
			res = (res * a) % _mod
		}
	}
	return res
}

func numberOfWays(startPos int, endPos int, k int) int {
	n := max(startPos, endPos) - min(startPos, endPos)
	if k < n || (k-n)&1 == 1 {
		return 0
	}
	n, k = k, n+(k-n)>>1 // c(n,k) = n!/k!/(n-k)!, 逆元吧

	c := _factorials[n]
	c = (c * binpow(_factorials[k], _mod-2)) % _mod
	c = (c * binpow(_factorials[n-k], _mod-2)) % _mod
	return c
}
