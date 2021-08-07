from typing import List

class Solution:
    def findKthNumber(self, m: int, n: int, k: int) -> int:
        l, r = 1, n*m
        while l < r:
            mid, cnt = l+r >> 1, 0
            for v in (min(mid//i, m) for i in range(1, min(mid, n) + 1)):
                cnt += v
            if cnt < k:
                l = mid + 1
            else:
                r = mid
        return l