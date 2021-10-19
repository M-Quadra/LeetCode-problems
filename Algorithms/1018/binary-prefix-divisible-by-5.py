from typing import List

class Solution:
    def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
        v = 0
        opt: List[bool] = []
        for n in nums:
            v = (v << 1 | n) % 5
            opt.append(v == 0)
        return opt
