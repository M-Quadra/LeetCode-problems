class Solution {
    func longestCommonSubsequence(_ text1: String, _ text2: String) -> Int {
        var dp = Array(repeating: Array(repeating: 0, count: text1.count+1), count: 2)
        var now = 0, pre = 1
        for ele in text2 {
            for it in text1.enumerated() {
                dp[now][it.offset+1] = dp[pre][it.offset] + (ele==it.element ? 1 : 0)
                dp[now][it.offset+1] = max(dp[now][it.offset+1], dp[now][it.offset], dp[pre][it.offset+1])
            }
            swap(&now, &pre)
        }
        return dp[pre].last ?? 0
    }
}