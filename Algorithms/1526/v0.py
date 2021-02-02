from typing import List

class Solution:
    def merge(self, sub: List[int]) -> (int, [int]):
        if len(sub) <= 2:
            return max(sub), []

        maxN, minN = max(sub), min(sub)
        if sub[1] != minN and sub[1] != maxN:
            return 0, [sub[0], sub[2]]
        
        if sub[1] == maxN:
            midN = max(sub[0], sub[2])
            return maxN-midN, [sub[0], sub[2]]
        
        midN = min(sub[0], sub[2])
        return midN-minN, [sub[0], sub[2]]

    def minNumberOperations(self, target: List[int]) -> int:
        if len(target) <= 2:
            opt, _ = self.merge(target)
            return opt
        
        opt, subAry = 0, target[:2]
        for v in target[2:]:
            subAry.append(v)
            v, subAry = self.merge(subAry)
            opt += v
        v, _ = self.merge(subAry)
        return opt + v