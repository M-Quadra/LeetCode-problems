class Solution {
public:
    bool canMeasureWater(int x, int y, int z) {
        if (x+y < z) return false;
        if (!x || !y) return x == z || y == z;
        
        if (x < y) std::swap(x, y);
        x %= y;
        z %= y;
        for (int i = 0; i < y; ++i) {
            if ((i*x)%y == z) return true;
        }
        return false;
    }
};