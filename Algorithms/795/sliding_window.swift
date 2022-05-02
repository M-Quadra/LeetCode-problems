class Solution {
    func numSubarrayBoundedMax(_ nums: [Int], _ left: Int, _ right: Int) -> Int {
        var j = 1, m = 1, opt = 0
        for i in nums.indices {
            j = max(j, i + 1)
            while nums.indices ~= j && nums[j] <= right {
                j += 1
            }
            
            switch nums[i] {
            case ..<left:
                m += 1
            case left...right:
                opt += m * (j - i)
                fallthrough
            default:
                m = 1
            }
        }
        return opt
    }
}