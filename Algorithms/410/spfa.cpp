struct step {
    int st, m, maxn;
    bool operator<(const step &b) const {
        return maxn > b.maxn;
    }
};

class Solution {
public:
    int splitArray(vector<int>& nums, int m) {
        nums.insert(nums.begin(), 0);
        for (auto it=nums.begin()+1; it!=nums.end(); ++it) *it += *(it-1);
        
        std::map<std::pair<int, int>, int> maxMap;
        maxMap[{nums.size(), m+1}] = nums.back();

        int avg = nums.back()/m;
        std::priority_queue<step> q; q.push({1, 1, 0});
        
        while (!q.empty()) {
            auto now = q.top(); q.pop();
            if (maxMap.count({now.st, now.m})) {
                if (now.maxn>=maxMap[{nums.size(), m+1}]) break;
                if (now.maxn>maxMap[{now.st, now.m}]) continue;
            }
            
            if (now.m == m) {
                int nxtVal = std::max(now.maxn, nums.back()-nums[now.st-1]);
                if (maxMap.count({nums.size(), m+1})) {
                    if (nxtVal >= maxMap[{nums.size(), m+1}]) continue;
                }
                
                maxMap[{nums.size(), m+1}] = nxtVal;
                continue;
            }
            
            // 回溯
            int v = std::max(nums[now.st-1] + now.maxn+1, nums[now.st-1] + avg);
            if (v > nums[now.st-1] + avg && now.st != 1) {
                q.push({1, 1, now.maxn});
            }
            
            // [st, ed]
            int maxEd = int(nums.size()-m+now.m-1);
            auto it = std::upper_bound(nums.begin()+now.st, nums.begin()+maxEd+1, v);
            int ed = std::min(int(it - nums.begin()), maxEd);
            int nxtVal = std::max(now.maxn, nums[ed]-nums[now.st-1]);
            if (!maxMap.count({ed+1, now.m+1}) || nxtVal<maxMap[{ed+1, now.m+1}]) {
                maxMap[{ed+1, now.m+1}] = nxtVal;
                q.push({ed+1, now.m+1, nxtVal});
            }
            
            if (now.st <= --ed) {
                nxtVal = std::max(now.maxn, nums[ed]-nums[now.st-1]);
                if (!maxMap.count({ed+1, now.m+1}) || nxtVal<maxMap[{ed+1, now.m+1}]) {
                    maxMap[{ed+1, now.m+1}] = nxtVal;
                    q.push({ed+1, now.m+1, nxtVal});
                }
            }
        }
        
        return maxMap[{nums.size(), m+1}];
    }
};