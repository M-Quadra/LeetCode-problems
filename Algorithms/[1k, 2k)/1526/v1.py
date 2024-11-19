from typing import List

class Solution:
    def minNumberOperations(self, target: List[int]) -> int:
        if len(target) <= 0:
            return 0
        
        opt = max(target[0], target[-1])
        for i in range(1, len(target)-1):
            if (target[0]-target[i])*(target[i]-target[i+1]) >= 0:
                continue
            if target[0]-target[i] < 0:
                opt += target[i]-max(target[0], target[i+1])
            else:
                opt += min(target[0], target[i+1]) - target[i]
        return opt