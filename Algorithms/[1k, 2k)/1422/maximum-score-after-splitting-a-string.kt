class Solution {
    fun maxScore(s: String): Int {
        var opt = 0
        val cnts = arrayOf(0, s.count { it == '1' })
        for (char in s.substring(0, s.lastIndex)) {
            if (char == '1') { cnts[1] -= 1 }
            else { cnts[0] += 1 }
            opt = maxOf(opt, cnts.sum())
        }
        return opt
    }
}