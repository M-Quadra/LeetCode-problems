import kotlin.math.max

class Solution {

    val prevs = IntArray(10_0000)
    val gens = IntArray(10_0000)

    fun totalSteps_0(nums: IntArray): Int {
        var opt = 0
        for ((i, v) in nums.withIndex()) {
            var (j, g) = listOf(i - 1, 1)
            while (j >= 0 && (gens[j] < g || nums[j] <= v)) {
                g = max(g, gens[j] + 1)
                j = prevs[j]
            }
            gens[i] = if (j >= 0) g else nums.size
            prevs[i] = j
            opt = max(opt, if (j >= 0) g else 0)
        }
        return opt
    }
}