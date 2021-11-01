class Solution:
    def hasAlternatingBits(self, n: int) -> bool:
        v = n&0b11
        if v==0b00 or v==0b11: return False
        while n > 0:
            if n&0b11 != v: return False
            n >>= 2
        return True