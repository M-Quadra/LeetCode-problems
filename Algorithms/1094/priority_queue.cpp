class Solution {
private:
    static bool comp(const std::vector<int> &a, const std::vector<int> &b) {
        return a[1]>b[1];
    }
public:
    bool carPooling(vector<vector<int>>& trips, int capacity) {
        std::priority_queue<std::vector<int>, std::vector<std::vector<int>>, decltype(&comp)> q(trips.begin(), trips.end(), comp);
        while (!q.empty()) {
            auto a = q.top(); q.pop();
            if (a[0] > capacity) return false;
            if (q.empty()) break;
            
            auto b = q.top();
            if (a[2] <= b[1]) continue;
            q.pop();
            
            q.push({a[0]+b[0], b[1], std::min(a[2], b[2])});
            if (a[2] > b[2]) q.push({a[0], b[2], a[2]});
            else if (a[2] < b[2]) q.push({b[0], a[2], b[2]});
        }
        return true;
    }
};