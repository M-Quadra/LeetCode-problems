class Solution {
    func bitwiseComplement(_ n: Int) -> Int {
        var a = 0, b = n
        while b > 0 {
            a = (-b & b)<<1
            b -= -b & b
        }
        return ~n & max(a-1, 1)
    }
}
