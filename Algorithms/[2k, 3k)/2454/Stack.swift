class Solution { // 单调栈
    func secondGreaterElement(_ nums: borrowing [Int]) -> [Int] {
        var ans = [Int](repeating: -1, count: nums.count)
        var stack0 = [Int](), stack1 = [Int]()
        for i in nums.indices {
            while let j = stack1.last, nums[j] < nums[i] { ans[stack1.removeLast()] = nums[i] }
            let st = stack1.count
            while let j = stack0.last, nums[j] < nums[i] { stack1.append(stack0.removeLast()) }
            stack0.append(i)
            stack1[st..<stack1.count].reverse()
        }
        return ans
    }
}
