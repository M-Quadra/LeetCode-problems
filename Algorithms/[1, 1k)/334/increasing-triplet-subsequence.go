package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	dp := [3]int{0x7FFFFFFF, 0x7FFFFFFF, 0x7FFFFFFF}

	for _, num := range nums {
		for i, v := range dp {
			if v < num {
				continue
			}
			if i == 2 {
				return true
			}
			dp[i] = num
			break
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
}
