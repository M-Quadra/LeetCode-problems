static auto _ = [](){
    std::ios::sync_with_stdio(false);
    std::cin.tie(NULL);
    return nullptr;
}();

class Solution {
public:
    int data[15][15];
    int numWays(int n, vector<vector<int>>& relation, int k) {
        memset(data, 0, sizeof data);
        data[0][0] = 1;
        for (int i = 0; i < k; ++i) {
            for (auto edge: relation) {
                if (!data[i][edge[0]]) continue;
                
                data[i+1][edge[1]] += data[i][edge[0]];
            }
        }
        return data[k][n-1];
    }
};