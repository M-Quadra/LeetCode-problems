class Solution {
public:
    struct pit {
        int x, y;
        
        bool operator<(const pit &b) const {
            if (x != b.x) return x < b.x;
            return y < b.y;
        }
    };
    
    vector<double> intersection(vector<int>& start1, vector<int>& end1, vector<int>& start2, vector<int>& end2) {
        if (std::max(start1[0], end1[0]) < std::min(start2[0], end2[0])) return {};
        if (std::max(start1[1], end1[1]) < std::min(start2[1], end2[1])) return {};
        
        int a1 = start1[1] - end1[1], b1 = end1[0] - start1[0];
        int c1 = -(a1*end1[0] + b1*end1[1]);
        
        int v1 = a1*start2[0] + b1*start2[1] + c1;
        int v2 = a1*end2[0] + b1*end2[1] + c1;
        if (v1*v2 > 0) return {};
        
        if (v1 == v2) {
            pit ary[4] = {
                {start1[0], start1[1]}, {end1[0], end1[1]},
                {start2[0], start2[1]}, {end2[0], end2[1]},
            };
            std::sort(ary, ary + 4);
            return {double(ary[1].x), double(ary[1].y)};
        }
        
        int a2 = start2[1] - end2[1], b2 = end2[0] - start2[0];
        int c2 = -(a2*end2[0] + b2*end2[1]);
        
        v1 = a2*start1[0] + b2*start1[1] + c2;
        v2 = a2*end1[0] + b2*end1[1] + c2;
        if (v1*v2 > 0) return {};
        
        double detL = a1*b2 - a2*b1;
        double L[2][2] = {
            {b2/detL, -a2/detL},
            {-b1/detL, a1/detL},
        };
        double optX = -c1*L[0][0] + -c2*L[1][0];
        double optY = -c1*L[0][1] + -c2*L[1][1];
        return {optX, optY};
    }
};