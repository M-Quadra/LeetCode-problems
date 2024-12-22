func distinctAverages(nums []int) int {
	set := make(map[int]struct{}, len(nums)/2)
	slices.Sort(nums)
	for i := 0; i*2 < len(nums); i++ {
		j := len(nums) - 1 - i
		set[nums[j]+nums[i]] = struct{}{}
	}
	return len(set)
}
