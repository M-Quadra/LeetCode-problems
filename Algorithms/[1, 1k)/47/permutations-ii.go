func reverseSlice(x interface{}, first, last int) {
	last--
	for swap := reflect.Swapper(x); first < last; first, last = first+1, last-1 {
		swap(first, last)
	}
}

func nextPermutationSlice(
	x interface{}, first, last int,
	less func(i, j int) bool,
) bool {
	i := last - 1
	if first >= last || first == i {
		return false
	}

	for swap := reflect.Swapper(x); ; {
		ip1 := i
		if i--; less(i, ip1) {
			j := last - 1
			for ; !less(i, j); j-- {
			}
			swap(i, j)
			reverseSlice(x, ip1, last)
			return true
		}
		if i == first {
			reverseSlice(x, first, last)
			return false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	opt := [][]int{}
	for sort.Ints(nums); ; {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		opt = append(opt, numsCopy)

		if !nextPermutationSlice(nums, 0, len(nums), func(i, j int) bool {
			return nums[i] < nums[j]
		}) {
			break
		}
	}
	return opt
}