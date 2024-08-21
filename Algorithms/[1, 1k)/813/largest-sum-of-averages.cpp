class Solution {
public:
    double largestSumOfAverages(vector<int>& A, int K) {
        double dp[105][105];
        memset(dp, 0, sizeof dp);
        
        int sum = 0, cnt = 0, len = (int)A.size() - K + 1;
        for (int i = 0; i < len; ++i) {
            sum += A[i];
            dp[0][i] = (double)sum/++cnt;
        }
        
        for (int ik = 1; ik < K; ++ik) {
            len = (int)A.size() - K + ik + 1;
            for (int st = ik; st < len; ++st) {
                sum = cnt = 0;
                for (int ed = st; ed < len; ++ed) {
                    sum += A[ed];
                    auto avg = (double)sum/++cnt;
                    dp[ik][ed] = max(dp[ik][ed], dp[ik-1][st-1] + avg);
                }
            }
        }
        
        return dp[K-1][A.size()-1];
    }
};