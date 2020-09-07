class Solution {
public:
    bool data[1005];
    bool canVisitAllRooms(vector<vector<int>>& rooms) {
        memset(data, 0, sizeof data);
        std::queue<int> q; q.push(0);
        while (!q.empty()) {
            int i = q.front(); q.pop();
            if (data[i]) continue;
            
            data[i] = true;
            for (auto nxt : rooms[i]) q.push(nxt);
        }
        
        for (int i = 0; i < rooms.size(); ++i) {
            if (!data[i]) return false;
        }
        return true;
    }
};