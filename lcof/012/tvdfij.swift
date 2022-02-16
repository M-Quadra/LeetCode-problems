class Solution {
    func pivotIndex(_ nums: [Int]) -> Int {
        var s = nums.suffix(from: 1).reduce(into: 0) { partialResult, v in
            partialResult += v
        }
        
        var v = 0
        for i in 0..<nums.count {
            if v == s {
                return i
            }
            
            v += nums[i]
            s -= i+1<nums.count ? nums[i+1] : 0
        }
        return -1
    }
}