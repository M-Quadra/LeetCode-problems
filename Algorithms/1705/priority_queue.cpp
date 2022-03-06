struct apple {
    int *num, day;
    friend bool operator<(apple a, const apple &b) {
        return b.day < a.day;
    }
};

class Solution {
public:
    int eatenApples(std::vector<int>& apples, std::vector<int>& days) {
        std::priority_queue<apple> q;
        int opt = 0;
        for (int i = 0; !q.empty() || i < apples.size(); ++i) {
            while (!q.empty() && (q.top().day<=i || *(q.top().num)<=0)) q.pop();
            if (i < apples.size() && days[i] > 0) {
                q.push({&apples[i], i + days[i]});
            }
            if (q.empty()) continue;
            
            -- *(q.top().num);
            ++opt;
        }
        
        return opt;
    }
};