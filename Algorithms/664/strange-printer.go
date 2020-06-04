package main

import "fmt"

var (
	dpAry   = [105][105]int{}
	dataAry = []int{}
)

// [st, ed]
func dp(st, ed int) int {
	if st > ed {
		return 0
	}
	if dpAry[st][ed] > 0 {
		return dpAry[st][ed]
	}

	cnt := 1 + dp(st+1, ed)
	for i := st + 1; i <= ed; i++ {
		if dataAry[i] != dataAry[st] {
			continue
		}

		cnt2 := dp(st, i-1) + dp(i+1, ed)
		if cnt2 < cnt {
			cnt = cnt2
		}
	}

	dpAry[st][ed] = cnt
	return cnt
}

func strangePrinter(s string) int {
	if len(s) <= 0 {
		return 0
	}

	dataAry = []int{int(s[0])}
	for i := 1; i < len(s); i++ {
		now := int(s[i])
		last := dataAry[len(dataAry)-1]
		if now == last {
			continue
		}

		dataAry = append(dataAry, now)
	}

	dpAry = [105][105]int{}
	return dp(0, len(dataAry)-1)
}

func main() {
	ts := "aaabbb"
	fmt.Println(strangePrinter(ts))
	return
}
