class Solution {
    
    func maxWidthRamp(_ nums: [Int]) -> Int {
        var arr = Array(repeating: Int.max, count: 5_0000 + 1)
        for (i, v) in nums.enumerated() {
            arr[v] = min(arr[v], i)
        }
        for i in 1..<arr.count {
            arr[i] = min(arr[i], arr[i-1])
        }
        var opt = 0
        for (i, v) in nums.enumerated().reversed() {
            opt = max(opt, i - arr[v])
        }
        return opt
    }
}