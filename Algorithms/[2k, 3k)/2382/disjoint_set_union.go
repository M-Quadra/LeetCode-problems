package main

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	dsu := make([]int64, len(nums))
	var add func(i int, v int64) int
	add = func(i int, v int64) int {
		dsu[i] += v
		if j := i + 1; j < len(nums) && nums[j] < 0 {
			nums[i] = -add(-nums[j], v)
		} else {
			nums[i] = -i
		}
		return -nums[i]
	}

	res, iMax := make([]int64, len(nums)), 0
	for i, q, isPrev := len(removeQueries)-1, 0, false; i >= 0; i-- {
		res[i], q = dsu[iMax], removeQueries[i]
		v, j := int64(nums[removeQueries[i]]), q-1
		if isPrev = 0 <= j && dsu[j] > 0; isPrev {
			v += dsu[j]
		}
		if q = add(q, v); dsu[q] > dsu[iMax] {
			iMax = q
		}
		if isPrev {
			nums[j] = -q
		}
	}
	return res
}
