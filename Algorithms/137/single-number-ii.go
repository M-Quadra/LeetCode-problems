package main

import (
	"fmt"
	"sort"
)

func main() {
	ary := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}

	fmt.Println(singleNumber(ary))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		lst := v - 1
		if i-1 > 0 {
			lst = nums[i-1]
		}

		nxt := v + 1
		if i+1 < len(nums) {
			nxt = nums[i+1]
		}

		if lst != v && nxt != v {
			return v
		}
	}
	return 0
}
