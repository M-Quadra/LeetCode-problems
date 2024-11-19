from typing import List
import collections

class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        opt = 0
        for k, v in collections.Counter(nums).items():
            if v > 1: continue
            opt += k
        return opt