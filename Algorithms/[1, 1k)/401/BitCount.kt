class Solution {
    fun readBinaryWatch(turnedOn: Int): List<String> {
        val opt = mutableListOf<String>()
        for (h in 0..11) {
            for (m in 0..59) {
                if (Integer.bitCount(h) + Integer.bitCount(m) != turnedOn) continue
                opt.add("$h:${if (m<10) "0$m" else m}")
            }
        }
        return opt
    }
}