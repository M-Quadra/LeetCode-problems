from typing import List

class Solution:
    def minNumberOperations(self, target: List[int]) -> int:
        if len(target) <= 0:
            return 0
        
        opt = target[0]
        for i in range(1, len(target)):
            opt += max(target[i]-target[i-1], 0)
        return opt