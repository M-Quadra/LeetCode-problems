map<int, int> cntMap;
bool cmp(int a, int b) {
    return cntMap[a] > cntMap[b];
}

class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        cntMap.clear();
        for (auto num : nums) cntMap[num] += 1;
        
        sort(nums.begin(), nums.end());
        nums.erase(unique(nums.begin(), nums.end()), nums.end());
        
        sort(nums.begin(), nums.end(), cmp);
        while (nums.size() > k) nums.pop_back();
        return nums;
    }
};