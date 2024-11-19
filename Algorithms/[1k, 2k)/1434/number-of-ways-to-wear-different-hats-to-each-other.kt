class Solution {
    
    fun numberWays(hats: List<List<Int>>): Int {
        val h2pArr = Array(41) { 0 }
        for ((i, v) in hats.withIndex()) {
            val p = 1 shl i
            for (h in v) {
                h2pArr[h] = h2pArr[h] or p
            }
        }

        val mod = 10_0000_0000 + 7
        val dpArr = Array(1 shl hats.size) { 0 }
        dpArr[0] = 1

        for (ps in h2pArr.filter { it > 0 }) {
            for ((i, num) in dpArr.withIndex().reversed()) {
                var ps = if (i or ps <= i || num <= 0) continue else ps
                do {
                    val p = -ps and ps
                    ps -= p
                    val nP = i or p
                    dpArr[nP] = if (nP <= i) continue else (dpArr[nP] + num) % mod
                } while (ps > 0)
            }
        }

        return dpArr.last()
    }
}