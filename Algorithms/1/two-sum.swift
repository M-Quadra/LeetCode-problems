class Solution {
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        let ary = nums.sorted()
        var l = 0
        var r = ary.count - 1
        
        while l < r {
            let tmp = ary[l] + ary[r]
            if tmp == target {
                break
            }
            
            if tmp < target {
                l += 1
            } else {
                r -= 1
            }
        }
        
        guard let lft = nums.firstIndex(of: ary[l]) else {
            return []
        }
        guard let rit = nums.lastIndex(of: ary[r]) else {
            return []
        }
        
        return [lft, rit]
    }
}
