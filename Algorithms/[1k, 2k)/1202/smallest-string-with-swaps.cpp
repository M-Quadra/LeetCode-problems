static auto _ = [](){
    std::ios::sync_with_stdio(false);
    std::cin.tie(NULL);
    return nullptr;
}();

class Solution {
public:
    int idx[100005];
    unsigned short data[100005][26];

    int find(int i) {
        return idx[i] = idx[i] == i ? i : find(idx[i]);
    }
    void unin(int a, int b) {
        int v = std::min(find(a), find(b));
        idx[idx[a]] = idx[idx[b]] = v;
        idx[a] = idx[b] = v;
    }
    
    string smallestStringWithSwaps(string s, vector<vector<int>>& pairs) {
        for (int i = 0; i < s.length(); ++i) idx[i] = i;
        for (auto p: pairs) unin(p[0], p[1]);
        pairs.clear();
        
        memset(data, 0, sizeof data);
        for (int i = 0; i < s.length(); ++i) ++data[find(i)][s[i] - 'a'];
        for (int i = 0; i < s.length(); ++i) {
            for (int v = idx[i], j = 0; j < 26; ++j) {
                if (data[v][j] <= 0) continue;
                
                --data[v][j];
                s[i] = 'a' + j;
                break;
            }
        }
        
        return s;
    }
};