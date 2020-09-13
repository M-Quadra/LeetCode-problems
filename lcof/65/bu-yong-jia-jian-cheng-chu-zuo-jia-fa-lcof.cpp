class Solution {
public:
    int add(int a, int b) {
        int opt = 0, f = 0;
        for (int i = 0; i < 32; ++i) {
            int v = (a&1) ^ (b&1);
            opt |= (v^f) << i;
            f = (a&b&1) | (v&f);
            a >>= 1; b >>=1;
        }
        return opt;
    }
};