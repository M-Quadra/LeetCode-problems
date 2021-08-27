from typing import List

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        opt: List[List[int]] = [[]]
        for v in nums:
            opt += [x+[v] for x in opt]
        return opt