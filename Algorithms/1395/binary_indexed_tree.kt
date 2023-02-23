class Solution {
    
    private var size = 0
    private val tre = Array(5) { IntArray(1000 + 5) }

    private fun add(y: Int, x: Int, v: Int) {
        var ix = if (v > 0) x else return
        while (ix <= size) {
            tre[y][ix] += v
            ix += -ix and ix
        }
    }

    private fun sum(y: Int, x: Int): Int {
        var (opt, ix) = arrayOf(0, x)
        while (ix > 0) {
            opt += tre[y][ix]
            ix -= -ix and ix
        }
        return opt
    }

    fun numTeams(rating: IntArray): Int {
        size = rating.size
        for (i in this.tre.indices) tre[i].fill(0, 0, size + 1)
        rating.withIndex()
            .sortedBy { it.value }
            .map { it.index }
            .withIndex()
            .forEach { (v, i) ->
                rating[i] = v + 1
            }

        for (v in rating) {
            add(0, v, 1)

            add(1, v, sum(0, v - 1))
            add(2, v, sum(1, v - 1))

            add(3, v, sum(0, size) - sum(0, v))
            add(4, v, sum(3, size) - sum(3, v))
        }
        return sum(2, size) + sum(4, size)
    }
}