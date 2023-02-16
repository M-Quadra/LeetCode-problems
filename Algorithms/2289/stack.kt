import kotlin.math.max

class Solution {

    val stack = Stack<Pair<Int, Int>>()

    fun totalSteps(nums: IntArray): Int {
        stack.clear()
        var opt = 0
        for (n in nums) {
            var g = 1
            while (!stack.empty()) {
                val top = stack.peek()
                if (top.first > n && top.second >= g) break
                g = max(g, stack.pop().second + 1)
            }
            stack.push(Pair(n, if (stack.isEmpty()) nums.size else g))
            opt = max(opt, if (stack.size == 1) 0 else g)
        }
        return opt
    }
}