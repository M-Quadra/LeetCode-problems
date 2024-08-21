package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	minV := A[0]
	maxV := A[len(A)-1]
	return max(maxV-minV-2*K, 0)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []int{2, 7, 2}
	fmt.Println(smallestRangeI(ts, 1))
}
