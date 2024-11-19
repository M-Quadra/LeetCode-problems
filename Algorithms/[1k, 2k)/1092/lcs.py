class Solution:
    def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
        dp, i = [["" for _ in range(len(str1)+1)] for _ in range(2)], 0
        for c2 in str2:
            i ^= 1
            for j, c1 in enumerate(str1, 1):
                if c1 == c2:
                    dp[i][j] = dp[i^1][j-1] + c1
                else:
                    dp[i][j] = dp[i][j-1] if len(dp[i][j-1]) > len(dp[i^1][j]) else dp[i^1][j]
        lcs = dp[i][len(str1)]
        
        opt = ""
        for v in lcs:
            i1, i2 = str1.find(v), str2.find(v)
            opt += str1[:i1] + str2[:i2] + v
            str1, str2 = str1[i1+1:], str2[i2+1:]
        return opt + str1 + str2