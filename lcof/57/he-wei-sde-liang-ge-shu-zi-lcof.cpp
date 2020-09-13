class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        for (int l = 0, r = (int)nums.size()-1; l < r;) {
            int v = nums[l] + nums[r];
            if (v == target) return {nums[l], nums[r]};
            
            v < target ? ++l : --r;
        }
        return {};
    }
};