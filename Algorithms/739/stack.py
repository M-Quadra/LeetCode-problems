from typing import List, Tuple

class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack: List[Tuple[int, int]] = []
        for i in range(len(temperatures)-1, -1, -1):
            v = temperatures[i]
            while stack and stack[-1][0] <= v: stack.pop()
            temperatures[i] = stack[-1][1] - i if stack else 0
            stack.append((v, i))
        return temperatures