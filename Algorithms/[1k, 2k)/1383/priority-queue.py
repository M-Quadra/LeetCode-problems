from typing import List
from queue import PriorityQueue

class Solution:
    def maxPerformance(self, n: int, speed: List[int], efficiency: List[int], k: int) -> int:
        indexes = sorted([i for i in range(n)], key=lambda x: efficiency[x], reverse=True)
        priQue, sum, opt = PriorityQueue[int](), 0, 0
        for i in indexes:
            if priQue.qsize() >= k:
                sum -= priQue.get()
            sum += speed[i]
            priQue.put(speed[i])
            opt = max(opt, sum * efficiency[i])
        return opt % (10**9 + 7)