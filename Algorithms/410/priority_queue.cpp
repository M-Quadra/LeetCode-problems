class Solution {
public:
    int splitArray(vector<int>& nums, int m) {
        int sum = 0, minn = 0;
        for (auto v: nums) {
            sum += v;
            minn = std::max(minn, v);
        }
        minn = std::max(minn, sum/m);
        
        std::priority_queue<int, std::vector<int>, std::greater<int>> q;
        q.push(minn--);
        while (!q.empty()) {
            int v = q.top(); q.pop();
            if (v <= minn) continue;
            
            
            int sum = 0, cnt = 0;
            for (int i = 0; i < nums.size(); ++i) {
                sum += nums[i];
                if (sum <= v && nums.size()-i+cnt+1 > m) continue;
                ++cnt;
                
                q.push(sum);
                if (i == 0) break;
                
                int l = sum - nums[i-1];
                if (l > minn) q.push(sum - nums[i-1]);
                
                if (nums[i] > v) break;
                sum = nums[i];
            }
            if (sum <= v) ++cnt;
            if (cnt == m) return v;
            
            minn = std::max(minn, v);
        }
        
        return sum;
    }
};