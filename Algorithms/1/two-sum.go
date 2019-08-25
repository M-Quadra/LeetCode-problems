package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 75, 25}
	target := 100
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	ary := make([]int, len(nums))
	copy(ary, nums)
	sort.Ints(ary)
	lft, rit := 0, len(ary)-1

	for lft < rit {
		tmp := ary[lft] + ary[rit]
		if tmp == target {
			break
		}
		if tmp < target {
			lft++
		} else {
			rit--
		}
	}

	optLft, optRit := 0, 0
	for i, v := range nums {
		if v == ary[lft] || v == ary[rit] {
			optLft = i
			break
		}
	}
	for i, v := range nums[optLft+1:] {
		if v == ary[lft] || v == ary[rit] {
			optRit = optLft + 1 + i
			break
		}
	}

	return []int{optLft, optRit}
}
