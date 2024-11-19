class Solution {
    fun sumBase(n: Int, k: Int): Int {
        var opt = 0
        var n = n
        while (n > 0) {
            opt += n % k
            n /= k
        }
        return opt
    }
}