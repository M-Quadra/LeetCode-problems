class Solution {
    
    fun numTeams(rating: IntArray): Int {
        val dp = Array(rating.size) { Array(2) { IntArray(2) } }
        for ((i, v) in rating.withIndex()) {
            for (j in (i + 1) until rating.size) {
                if (rating[j] == v) continue
                val k = if (rating[j] < v) 0 else 1
                dp[j][0][k] += 1
                dp[j][1][k] += dp[i][0][k]
            }
        }

        var opt = 0
        for (i in rating.indices) opt += dp[i][1].sum()
        return opt
    }
}