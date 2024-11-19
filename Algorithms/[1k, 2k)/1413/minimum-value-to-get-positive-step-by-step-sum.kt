class Solution {
    fun minStartValue(nums: IntArray): Int {
        var minN = nums[0]
        for (i in 1..nums.lastIndex) {
            nums[i] += nums[i - 1]
            minN = minOf(minN, nums[i])
        }
        return maxOf(1, 1 - minN)
    }
}