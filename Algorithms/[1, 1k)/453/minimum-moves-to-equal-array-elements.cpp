class Solution {
public:
    int minMoves(vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        int opt = 0;
        for (int v : nums) {
            opt += v - nums.front();
        }
        return opt;
    }
};