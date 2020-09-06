class Solution {
public:
    vector<int> numMovesStonesII(vector<int>& stones) {
        std::sort(stones.begin(), stones.end());
        int sumGap = (int)stones.back() - (int)stones.front() - (int)stones.size() + 1;
        for (int i = (int)stones.size()-1; i > 0; --i) stones[i] -= stones[i-1] + 1;
        stones.erase(stones.begin());
        int maxNum = sumGap - std::min(stones.front(), stones.back());
        
        int minNum = maxNum;
        for (int l = 0, r = 0, gapNum = stones[0]; r < stones.size();) {
            sumGap = (int)stones.size() - (r - l + 1);
            if (sumGap <= 0) {
                gapNum -= stones[l++];
                continue;
            }
            
            if (sumGap < gapNum) {
                minNum = std::min(minNum, sumGap + 1);
                gapNum -= stones[l++];
                if (l > r) gapNum += stones[++r];
            } else if (sumGap == gapNum) {
                minNum = std::min(minNum, sumGap);
                gapNum -= stones[l++];
                gapNum += stones[++r];
            } else {
                gapNum += stones[++r];
            }
        }
        
        return {minNum, maxNum};
    }
};