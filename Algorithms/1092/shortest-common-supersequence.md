# LeetCode 1092 最短公共超序列

贪心莽不出来, DP看起来像, 但我依然可耻地选择SPFA低空飞过

最优解应该是LCS(dp)

羸弱如我, 看到题解才恍然大悟

先求最长公共子序列, 再根据LCS合并两个字符串
