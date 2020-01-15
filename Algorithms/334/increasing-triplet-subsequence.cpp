class Solution {
public:
    bool increasingTriplet(vector<int>& nums) {
        if (nums.size() < 3) return false;
        
        int dp[3];
        memset(dp, 0x7F, sizeof dp);
        
        for (int num : nums) {
            for (int i = 0; i < 3; ++i) {
                if (dp[i] < num) continue;
                if (i == 2) return true;
                dp[i] = num;
                break;
            }
        }
        
        return false;
    }
};