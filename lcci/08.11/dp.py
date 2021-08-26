maxLen = 10**6 + 5
dp = [1 for _ in range(maxLen)]
for c in [5, 10, 25]:
    for i in range(c, maxLen, 1):
        dp[i] += dp[i-c]
dp[0] = 0

class Solution:
    def waysToChange(self, n: int) -> int:
        return dp[n] % 1000000007