class Solution:
    def integerReplacement(self, n: int) -> int:
        opt = 0
        while n > 1:
            if n > 3 and n & 3 == 3:
                n += 1
            elif n & 1 == 1: 
                n -= 1
            else:
                n >>= 1
            opt += 1
        return opt