class Solution {

    fun reverseBits(num: Int): Int {
        var (opt, num, cnt0, cnt1) = listOf(1, num, 0, 0)
        while (num != 0) {
            while (num.and(1) == 0) {
                cnt0 += 1
                num = num.ushr(1)
            }

            var tmp = 0
            while (num.and(1) == 1) {
                tmp += 1
                num = num.ushr(1)
            }

            opt = maxOf(opt, tmp + 1 + if (cnt0 <= 1) cnt1 else 0)
            cnt1 = tmp
            cnt0 = 0
        }

        return minOf(opt, 32)
    }
}