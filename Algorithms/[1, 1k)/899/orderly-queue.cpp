class Solution {
public:
    string orderlyQueue(string S, int K) {
        if (K > 1) {
            std::sort(S.begin(), S.end());
            return S;
        }
        
        auto minS = S;
        for (int i = 0; i < S.length(); ++i) {
            if (S < minS) minS = S;
            S.append({S.front()});
            S.erase(S.begin());
        }
        
        return minS;
    }
};