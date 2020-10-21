class Solution {
public:
    vector<int> drawLine(int length, int w, int x1, int x2, int y) {
        auto opt = std::vector<int>(length);
        int lNum = w / 32;
        for (int i = x1; i <= x2; ++i) {
            int idx = y*lNum + i/32;
            opt[idx] |= 1 << (31 - (i%32));
        }
        
        return opt;
    };
};