package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	cnt := len(stones)
	if cnt < 2 {
		return stones[0]
	}

	for {
		sort.Ints(stones)
		s1 := stones[cnt-1]
		s2 := stones[cnt-2]
		if s1*s2 <= 0 {
			break
		}
		stones[cnt-1] = s1 - s2
		stones[cnt-2] = 0
	}

	return stones[cnt-1]
}

func main() {
	ts := lastStoneWeight([]int{1, 3})
	fmt.Println(ts)
}
