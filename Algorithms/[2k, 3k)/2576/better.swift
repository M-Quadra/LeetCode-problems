struct Solution: ~Copyable {
    consuming func maxNumOfMarkedIndices(_ nums: consuming [Int]) -> Int {
        var nums = nums.sorted(), len = 0, r = nums.count-1
        for l in (0...nums.count/2).reversed() {
            guard l+len < r, nums[l]*2 <= nums[r] else { continue }
            len += 1
            r -= 1
        }
        return len*2
    }
}
