class Solution {
    private fun nextPermutation(arr: Array<Int>): Boolean {
        if (arr.isEmpty()) return false
        val first = 0
        val last = arr.size

        var i = arr.lastIndex
        if (first >= last || first == i) return false

        while (true) {
            val ip1 = i--
            if (arr[i] < arr[ip1]) {
                var j = last - 1
                while (!(arr[i] < arr[j])) --j
                val tmp = arr[i]
                arr[i] = arr[j]
                arr[j] = tmp

                val tmpArr = arr.slice(ip1..last - 1).reversed()
                for (i in 0..tmpArr.lastIndex) {
                    arr[ip1 + i] = tmpArr[i]
                }
                return true
            }

            if (i == first) {
                arr.reverse()
                return false
            }
        }
    }

    fun readBinaryWatch(turnedOn: Int): List<String> {
        val arr = Array(11) { i ->
            if (i >= 11-turnedOn) 1 else 0
        }
        val opt = mutableListOf<String>()

        var lastH = -1
        var lastM = -1
        do {
            var h = 0
            for (i in 0 .. 3) {
                h = (h shl 1) or arr[i]
            }
            var m = 0
            for (i in 4 .. 10) {
                m = (m shl 1) or arr[i]
            }
            if (h == lastH && m == lastM) continue
            if (h > 11 || m > 59) continue
            lastH = h
            lastM = m
            opt.add(String.format("%d:%02d", h, m))
        } while (nextPermutation(arr))
        return opt
    }
}