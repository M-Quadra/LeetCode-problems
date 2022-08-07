class Solution {
    
    func longestConsecutive(_ nums: [Int]) -> Int {
        var set = Set(nums), opt = 0
        while let n = set.popFirst() {
            var prev = 0, next = 0
            while set.remove(n-1 - prev) != nil {
                prev += 1
            }
            while set.remove(n+1 + next) != nil {
                next += 1
            }
            opt = max(opt, 1 + prev + next)
        }
        return opt
    }
}