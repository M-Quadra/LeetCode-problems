var _facs = make([]int, 1e5+1)

const _mod int = 1e9 + 7

func init() {
	_facs[0] = 1
	for i := 1; i <= 1e5; i++ {
		_facs[i] = _facs[i-1] * i
		_facs[i] %= _mod
	}
}

func mult(v int, arr []int) int {
	sum := 0
	deno := 1
	for _, v := range arr {
		if v <= 0 {
			continue
		}
		sum += v
		deno = (deno * _facs[v]) % _mod
	}
	if sum <= 0 {
		return v
	} else if v == 0 {
		v = _facs[sum]
	} else {
		v = (v * _facs[sum]) % _mod
	}

	// 逆元
	inve := 1
	for a, b := deno, _mod-2; b != 0; b >>= 1 {
		if b&1 == 1 {
			inve = (a * inve) % _mod
		}
		a = (a * a) % _mod
	}
	return (v * inve) % _mod
}

func countAnagrams(s string) int {
	ans := 0

	arr := make([]int, 26)
	for _, char := range s {
		if char == ' ' {
			ans = mult(ans, arr)
			arr = make([]int, 26)
			continue
		}
		arr[char-'a']++
	}
	ans = mult(ans, arr)
	return ans
}
