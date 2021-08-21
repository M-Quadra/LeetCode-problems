from typing import List

class Solution:
    def kConcatenationMaxSum(self, arr: List[int], k: int) -> int:
        maxSub, tmpSub = 0, 0
        maxSuf, tmpSuf = 0, 0
        maxPre, s = 0, 0
        for i, v in enumerate(arr):
            tmpSub = max(0, v, tmpSub + v)
            maxSub = max(maxSub, tmpSub)

            tmpSuf += arr[-1 - i]
            maxSuf = max(maxSuf, tmpSuf)

            s += v
            maxPre = max(maxPre, s)
        if k <= 1:
            return maxSub
        s = max(0, s)

        opt = maxSuf + s * (k-2) + maxPre
        return max(opt, maxSub) % (10**9 + 7)