import math

class Solution:
    def waysToChange(self, n: int) -> int:
        if n <= 0:
            return 0

        n //= 5
        return math.floor((n+8) * (2*(n**2) + 11*n + 18) / 120) % 1000000007