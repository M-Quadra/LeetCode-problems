class Solution {
    fun minOperations(nums: IntArray): Int {
        var opt = 0
        for (i in 1..nums.lastIndex) {
            opt += maxOf(0, nums[i - 1] + 1 - nums[i])
            nums[i] = maxOf(nums[i], nums[i - 1] + 1)
        }
        return opt
    }
}