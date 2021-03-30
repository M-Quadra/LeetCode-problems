from typing import List
import bisect

class Solution:
    def numFriendRequests(self, ages: List[int]) -> int:
        ages = sorted([x for x in filter(lambda x: x>14, ages)])

        opt = 0
        lastV, lastS = -1, -1
        for i, v in enumerate(ages):
            if v == lastV:
                opt += lastS
                continue
            if i+1 == len(ages):
                continue

            lastV, lastS = v, 0
            ed = bisect.bisect_left(ages, 2*v - 14, i)
            if ed:
                lastS = ed - i - 1
                opt += lastS
        return opt