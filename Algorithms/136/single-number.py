from typing import List

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        opt = 0
        for v in nums:
            opt ^= v
        return opt