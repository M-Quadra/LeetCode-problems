from typing import List

class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        nums.sort()
        sumAry = nums.copy()
        for i in range(1, len(sumAry)):
            sumAry[i] += sumAry[i-1]
        nums.reverse()
        sumAry.reverse()

        opt: int = 0
        def dfs(i: int, v: int) -> None:
            if i == len(nums):
                nonlocal opt
                opt += 1 if v == target else 0
                return

            minN, maxN = v-sumAry[i], v+sumAry[i]
            if target < minN or maxN < target:
                return
            
            dfs(i+1, v-nums[i])
            dfs(i+1, v+nums[i])
        dfs(0, 0)
        return opt
