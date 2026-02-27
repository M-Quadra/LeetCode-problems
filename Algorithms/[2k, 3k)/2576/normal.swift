struct Solution: ~Copyable {
    consuming func maxNumOfMarkedIndices(_ nums: consuming [Int]) -> Int {
        var nums = nums.sorted(), ans = 0, r = nums.count-1
        for l in stride(from: nums.count/2, to: -1, by: -1) {
            nums[l].negate()
            if -nums[l]*2 <= nums[r] {
                ans += 1
                r -= 1
            } else if nums[r] < 0, -nums[l]*2 <= -nums[r] {
                nums[r].negate()
            } else {
                nums[l+ans].negate()
            }
        }
        return ans * 2
    }
}
