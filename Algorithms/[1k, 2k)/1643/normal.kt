private val cache = Array(30) { Array<Double?>(30) { null } }
private fun pai(st: Int, ed: Int): Double = cache[st][ed] ?: run {
    val v = if (st == ed) st.toDouble() else pai(st, ed - 1) * ed.toDouble()
    cache[st][ed] = v
    return@run v
}

class Solution {

    fun kthSmallestPath(destination: IntArray, k: Int): String {
        var currentK = 0
        return String(CharArray(destination.sum()) {
            val (v, h) = destination
            if (v <= 0) return@CharArray 'H'
            if (h <= 0) return@CharArray 'V'

            val n = (pai(h, v + h - 1) / pai(1, v)).toInt()
            if (n + currentK < k) {
                currentK += n
                destination[0] -= 1
                return@CharArray 'V'
            }

            destination[1] -= 1
            return@CharArray 'H'
        })
    }
}