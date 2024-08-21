class Solution {
    func increasingTriplet(_ nums: [Int]) -> Bool {
        if nums.count < 3 {
            return false
        }
        
        var dp = Array.init(repeating: 0x7FFFFFFF, count: 3)
        for num in nums {
            for i in 0..<dp.count {
                if dp[i] < num {
                    continue
                }
                if i == 2 {
                    return true
                }
                dp[i] = num
                break;
            }
        }
        
        return false
    }
}