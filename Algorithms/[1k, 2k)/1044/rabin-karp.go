const _mod = uint64(1<<61) - 1

func mul(a, b uint64) uint64 {
	l1, h1 := uint64(uint32(a)), uint64(a>>32)
	l2, h2 := uint64(uint32(b)), uint64(b>>32)
	l, m, h := l1*l2, l1*h2+l2*h1, h1*h2
	ret := (l & _mod) + (l >> 61) + (h << 3) + (m >> 29) + (m << 35 >> 3) + 1
	ret = (ret & _mod) + (ret >> 61)
	ret = (ret & _mod) + (ret >> 61)
	return ret - 1
}

var _mulArr = make([]uint64, 1e5+5)

func init() {
	_mulArr[0] = 1
	for i := 1; i < 1e5+5; i++ {
		_mulArr[i] = mul(_mulArr[i-1], 26)
	}
}

func longestDupSubstring(s string) string {
	opt := ""
	for l, r := 0, len(s)+1; l < r; {
		m := l + (r-l)/2
		h := uint64(0)
		for i := 0; i < m; i++ {
			h += mul(_mulArr[m-i-1], uint64(s[i]-'a'))
			h = mul(h, 1)
		}

		set := map[uint64]struct{}{h: {}}
		for i := m; i < len(s); i++ {
			h = mul(h, 26) + uint64(s[i]-'a')
			j := mul(_mulArr[m], uint64(s[i-m]-'a'))
			if h < j {
				h += _mod
			}
			h -= j
			h = mul(h, 1)
			if _, ok := set[h]; ok {
				opt = s[i+1-m : i+1]
				break
			}
			set[h] = struct{}{}
		}

		if len(opt) >= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return opt
}
