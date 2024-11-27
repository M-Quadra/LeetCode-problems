class Solution {
    func unequalTriplets(_ nums: borrowing [Int]) -> Int {
        var arr = [Int](repeating: 0, count: 1001)
        for i in nums.indices { arr[nums[i]] += 1 }
        
        var l = 0, cnt = 0
        for n in arr where n > 0 {
            cnt += l * n * (nums.count - l - n)
            l += n
        }
        return cnt
    }
}
