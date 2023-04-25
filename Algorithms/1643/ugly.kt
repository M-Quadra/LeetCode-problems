class Solution {
    
    fun kthSmallestPath(destination: IntArray, k: Int): String {
        var mark = 0
        var lastK = k - 1

        for (v in destination[0] downTo 1) {
            val (offset, n) = run {
                val keep = IntArray(2) { 1 }
                var n = 0
                for (h in 1..destination[1]) {
                    val nN = h * keep[0] / keep[1]
                    n = if (nN <= lastK) nN else return@run Pair(h - 1, n)
                    keep[0] *= (v + h)
                    if (keep[0] / (h + 1) * (h + 1) == keep[0]) {
                        keep[0] /= (h + 1)
                    } else {
                        keep[1] *= (h + 1)
                    }
                }
                return@run Pair(destination[1], n)
            }
            mark += 1 shl (v - 1 + offset)
            lastK -= n
        }

        return String(CharArray(destination.sum()) {
            val i = destination.sum() - it - 1
            if (mark.and(1 shl i) == 0) 'H' else 'V'
        })
    }
}