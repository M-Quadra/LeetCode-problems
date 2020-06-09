class Solution {
public:
    int dpAry[105][105];
    
    int strangePrinter(string s) {
        memset(dpAry, 0, sizeof dpAry);
        auto ed = unique(s.begin(), s.end());
        while (s.end() != ed) s.pop_back();
        if (s.length() <= 1) return (int)s.length();
        
        return dp(s, 0, s.length()-1);
    }
    
    // [st, ed]
    int dp(string &s, unsigned long st, unsigned long ed) {
        if (ed < st) return 0;
        if (dpAry[st][ed] > 0) return dpAry[st][ed];
        
        int len = 1 + dp(s, st+1, ed);
        for (auto i = st+1; i <= ed; ++i) {
            if (s[i] != s[st]) continue;
            len = min(len, dp(s, st, i-1) + dp(s, i+1, ed));
        }
        
        dpAry[st][ed] = len;
        return len;
    }
};
