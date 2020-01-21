class Solution {
    func largestSumOfAverages(_ A: [Int], _ K: Int) -> Double {
        var dp = Array.init(repeating: Array.init(repeating: 0.0, count: A.count), count: K)
        
        var sum = 0
        var cnt = 0
        var len = A.count - K + 0 + 1
        for i in 0..<len {
            sum += A[i]
            cnt += 1
            dp[0][i] = Double(sum)/Double(cnt)
        }
        
        for ik in 1..<K {
            len = A.count - K + ik + 1
            for st in ik..<len {// [st, ed]
                sum = 0
                cnt = 0
                for ed in st..<len {
                    sum += A[ed]
                    cnt += 1
                    let avg = Double(sum)/Double(cnt)
                    dp[ik][ed] = max(dp[ik][ed], dp[ik-1][st-1] + avg)
                }
            }
        }
        
        return dp[K - 1][A.count - 1]
    }
}