class Solution {
    func unequalTriplets(_ nums: borrowing [Int]) -> Int {
        var cnt = 0
        for i in nums.indices {
            for j in (i+1)..<nums.count where nums[i] != nums[j] {
                for k in (j+1)..<nums.count where nums[j] != nums[k] {
                    if nums[i] != nums[k] { cnt += 1 }
                }
            }
        }
        return cnt
    }
}
