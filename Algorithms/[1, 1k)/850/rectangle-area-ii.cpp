class Solution {
public:
    int rectangleArea(vector<vector<int>>& rectangles) {
        long long opt = 0;
        long long mod = 1000000007;
        
        while (rectangles.size()) {
            auto area = ((rectangles[0][2] - rectangles[0][0])%mod) * ((rectangles[0][3] - rectangles[0][1])%mod);
            opt += area % mod;
            opt %= mod;
            
            for (int i = 1, ed = (int)rectangles.size(); i < ed; ++i) {
                auto subRecAry = subRec(rectangles[0], rectangles[i]);
                if (subRecAry.size() > 0) {
                    rectangles[i] = subRecAry[0];
                    for (int j = 1; j < subRecAry.size(); ++j) {
                        rectangles.push_back(subRecAry[j]);
                    }
                }
            }
            
            rectangles.erase(rectangles.begin());
        }
        
        return (int)opt;
    }
    
    vector<int> subLine(vector<int> line1, vector<int> line2) {
        vector<int> optAry = {max(line1[0], line2[0]), min(line1[1], line2[1])};
        if (optAry[0] < optAry[1]) {
            return optAry;
        }
        
        return {};
    }
    
    vector<vector<int>> subRec(vector<int> rec1, vector<int> rec2) {
        auto x = subLine({rec1[0], rec1[2]}, {rec2[0], rec2[2]});
        if (x.size() <= 0) return {};
        auto y = subLine({rec1[1], rec1[3]}, {rec2[1], rec2[3]});
        if (y.size() <= 0) return {};
        
        vector<vector<int>> optAry;
        
        vector<int> rec = {rec2[0], rec2[1], x[0], rec2[3]};// left
        if (rec[0] < rec[2] && rec[1] < rec[3]) {
            optAry.push_back(rec);
        }
        
        rec = {x[0], y[1], x[1], rec2[3]};// up
        if (rec[0] < rec[2] && rec[1] < rec[3]) {
            optAry.push_back(rec);
        }
        
        rec = {x[0], rec2[1], x[1], y[0]};// down
        if (rec[0] < rec[2] && rec[1] < rec[3]) {
            optAry.push_back(rec);
        }
        
        rec = {x[1], rec2[1], rec2[2], rec2[3]};// right
        if (rec[0] < rec[2] && rec[1] < rec[3]) {
            optAry.push_back(rec);
        }
        
        if (optAry.size() <= 0) {
            optAry.push_back({0, 0, 0, 0});
        }
        return optAry;
    }
};