class Solution {
public:
    int scoreAry[75][75][75];
    
    int cherryPickup(vector<vector<int>>& grid) {
        memset(scoreAry, -1, sizeof scoreAry);
        
        int h = (int)grid.size();
        int w = (int)grid[0].size();
        
        std::queue<std::pair<int, std::pair<int, int>>> q;
        q.push({0, {0, w-1}});
        
        scoreAry[0][0][w-1] = grid[0][0] + grid[0][w-1];
        
        while (!q.empty()) {
            auto nowPit = q.front(); q.pop();
            int nowY = nowPit.first;
            
            if (nowY+1 >= h) continue;
            
            int nowL = nowPit.second.first, nowR = nowPit.second.second;
            int nowS = scoreAry[nowY][nowL][nowR];
            
            for (int i = -1; i <= 1; ++i) {
                int l = nowL + i;
                if (l < 0 || w <= l) continue;
                for (int j = -1; j <= 1; ++j) {
                    int r = nowR + j;
                    if (r < 0 || w <= r) continue;
                    
                    int nxtS = nowS + grid[nowY+1][l];
                    if (l != r) nxtS += grid[nowY+1][r];
                    if (nxtS <= scoreAry[nowY+1][l][r]) continue;
                    
                    scoreAry[nowY+1][l][r] = nxtS;
                    q.push({nowY+1, {l, r}});
                }
            }
        }
        
        int opt = 0;
        for (int i = 0; i < w; ++i) {
            for (int j = 0; j < w; ++j) {
                opt = std::max(opt, scoreAry[h-1][i][j]);
            }
        }
        return opt;
    }
};