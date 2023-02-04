class Solution {

    fun checkIfExist(arr: IntArray): Boolean {
        val map = mutableMapOf<Int, Int>()
        for (v in arr) {
            map[v] = (map[v] ?: 0) + 1
        }
        for (k in map.keys) {
            val n = map[k * 2] ?: continue
            if (n >= (if (k == 0) 2 else 1)) return true
        }
        return false
    }
}