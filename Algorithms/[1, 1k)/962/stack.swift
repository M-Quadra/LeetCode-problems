class Solution {
    
    func maxWidthRamp(_ nums: [Int]) -> Int {
        return (1..<nums.count).reduce(into: [0]) { stack, i in
            if nums[stack.last!] > nums[i] { stack.append(i) }
        }.reversed().reduce(into: (0, nums.count-1)) { it, i in
            while (it.1>0 && nums[i]>nums[it.1]) { it.1 -= 1 }
            it.0 = max(it.0, it.1 - i)
        }.0
    }
}