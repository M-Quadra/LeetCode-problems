class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        while (~--k) {
            for (int i = ((int)nums.size()-2)>>1; ~i; --i) {
                int l = i<<1|1, r = l+1;
                if (r+1 > nums.size()) {
                    if (nums[i] < nums[l]) swap(nums[i], nums[l]);
                    continue;
                }
                
                if (nums[i] < nums[l]) swap(nums[i], nums[l]);
                if (nums[i] < nums[r]) swap(nums[i], nums[r]);
            }
            
            if (!k) return nums[0];
            swap(nums[0], nums[nums.size()-1]);
            nums.pop_back();
        }
        return 0;
    }
};