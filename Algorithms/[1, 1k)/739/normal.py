from typing import List

class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        earliest = [0x7FFFFF for _ in range(105)]
        for i in range(len(temperatures)-1, -1, -1):
            v = temperatures[i]
            temperatures[i] = earliest[v+1] - i if earliest[v+1]<len(temperatures) else 0
            for j in range(v+1): earliest[j] = i
        return temperatures