class Solution {
public:
    char data[25][25];
    
    int dir[4][2] = {
        -1, 0,
        1, 0,
        0, -1,
        0, 1,
    };
    
    bool spfa(int stY, int stX, int edY, int edX, int boxY, int boxX) {
        if (stY == edY && stX == edX) return true;
        if (data[edY][edX] == -1) return false;
        
        char tmpData[25][25];
        memcpy(tmpData, data, sizeof data);
        tmpData[stY][stX] = -1;
        tmpData[boxY][boxX] = -1;
        tmpData[edY][edX] = 0;
        
        std::queue<std::pair<int, int>> q;
        q.push({stY, stX});
        while (!q.empty()) {
            int nowY = q.front().first;
            int nowX = q.front().second;
            q.pop();
            
            for (int i = 0; i < 4; ++i) {
                int nextY = nowY + dir[i][0];
                int nextX = nowX + dir[i][1];
                if (tmpData[nextY][nextX] == -1) continue;
                
                if (nextY == edY && nextX == edX) return true;
                
                tmpData[nextY][nextX] = -1;
                q.push({nextY, nextX});
            }
        }
        
        return false;
    }
    
    struct step {
        int boxY, boxX;
        int poiY, poiX;
        int s;
    };
    
    int minPushBox(vector<vector<char>>& grid) {
        int boxY = 0, boxX = 0;
        int stY = 0, stX = 0;
        int edY = 0, edX = 0;
        
        memset(data, -1, sizeof data);
        for (int y = 0; y < grid.size(); ++y) {
            for (int x = 0; x < grid[y].size(); ++x) {
                switch (grid[y][x]) {
                    case '.': data[y+1][x+1] = 0;     break;
                    case 'S': stY = y+1; stX = x+1;   break;
                    case 'T': edY = y+1; edX = x+1;   break;
                    case 'B': boxY = y+1; boxX = x+1; break;
                    default: break;
                }
            }
        }
        data[stY][stX] = 0;
        data[edY][edX] = 0;
        data[boxY][boxX] = 0;
        
        std::queue<step> q;
        q.push({boxY, boxX, stY, stX, 0});
        while (!q.empty()) {
            auto nowStep = q.front(); q.pop();

            for (int i = 0; i < 4; ++i) {
                int nextBoxY = nowStep.boxY + dir[i][0];
                int nextBoxX = nowStep.boxX + dir[i][1];
                int d = 1 << i;
                if (data[nextBoxY][nextBoxX] & d) continue;
                
                int edPoiY = nowStep.boxY + dir[i^1][0];
                int edPoiX = nowStep.boxX + dir[i^1][1];
                if (!spfa(nowStep.poiY, nowStep.poiX, edPoiY, edPoiX, nowStep.boxY, nowStep.boxX)) continue;
                
                if (nextBoxY == edY && nextBoxX == edX) return nowStep.s+1;
                
                data[nextBoxY][nextBoxX] |= d;
                q.push({nextBoxY, nextBoxX, nowStep.boxY, nowStep.boxX, nowStep.s+1});
            }
        }
        
        return -1;
    }
};