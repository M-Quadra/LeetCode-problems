package main

import "strconv"

func nextBeautifulNumber(n int) int {
	for n++; ; n++ {
		iToCnt := [10]int{}
		for _, v := range strconv.Itoa(n) {
			iToCnt[v-'0']++
		}
		for i, cnt := range iToCnt {
			if cnt > 0 && i != cnt {
				goto ed
			}
		}
		return n
	ed:
	}
}
