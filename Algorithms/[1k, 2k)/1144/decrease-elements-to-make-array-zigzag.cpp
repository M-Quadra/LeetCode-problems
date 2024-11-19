class Solution {
public:
    int movesToMakeZigzag(vector<int>& nums) {
        nums.insert(nums.begin(), 0x7FFFFF);
        nums.push_back(0x7FFFFF);
        
        int optOdd = 0, optEven = 0;
        for (int i = 1; i+1 < nums.size(); ++i) {
            int cnt = nums[i] - std::min(nums[i-1], nums[i+1]) + 1;
            cnt = std::max(0, cnt);
            if (i&1) {
                optOdd += cnt;
                continue;
            }
            
            optEven += cnt;
        }
        return std::min(optOdd, optEven);
    }
};