from typing import List

class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        dp = [0x7FFFFFFF for _ in range(len(cost)+2)]
        dp[0] = dp[1] = 0
        for i, v in enumerate(cost):
            v += dp[i]
            dp[i+1] = min(dp[i+1], v)
            dp[i+2] = min(dp[i+2], v)
        return dp[len(cost)]