class Solution {

    fun wiggleMaxLength(nums: IntArray): Int {
        val st = nums.indices.firstOrNull {
            it > 0 && nums[it - 1] != nums[it]
        } ?: return 1

        var (opt, dif) = arrayOf(2, nums[st] - nums[st - 1])
        for (i in st until nums.size) {
            dif = if ((nums[i] - nums[i - 1]) * dif >= 0) continue else -dif
            opt += 1
        }
        return opt
    }
}