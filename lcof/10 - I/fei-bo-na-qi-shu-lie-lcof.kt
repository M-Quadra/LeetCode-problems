private const val mod = 1000000007
private val arr = IntArray(101).apply {
    this[1] = 1
    for (i in 2..this.lastIndex) {
        this[i] = (this[i - 2] + this[i - 1]) % mod
    }
}

class Solution {

    fun fib(n: Int): Int {
        return arr[n]
    }
}