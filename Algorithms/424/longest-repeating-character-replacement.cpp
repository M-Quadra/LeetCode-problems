int cntAry[26];

class Solution {
public:
    int maxCnt() {
        int mc = 0;
        for (int i = 0; i < 26; ++i) {
            mc = max(mc, cntAry[i]);
        }
        return mc;
    }
    
    int characterReplacement(string s, int k) {
        if (s.length() <= 0) return 0;
        
        memset(cntAry, 0, sizeof cntAry);
        cntAry[s[0] - 'A'] = 1;
        
        int iL = 0, iR = 0;
        int maxLen = 1;
        
        while (iR + 1 < s.length()) {
            for (int i; iR + 1 < s.length(); ++iR) {
                i = s[iR + 1] - 'A';
                ++cntAry[i];
                if (iR - iL + 1 - maxCnt() < k) continue;
                
                --cntAry[i];
                break;
            }
            
            maxLen = max(maxLen, iR - iL + 1);
            if (iR >= s.length()) break;
            
            while (iL + 1 < s.length()) {
                --cntAry[s[iL] - 'A'];
                ++iL;
                if (s[iL - 1] != s[iL]) break;
            }
            
            while (iL > iR) ++cntAry[s[++iR] - 'A'];
        }
        
        return maxLen;
    }
};