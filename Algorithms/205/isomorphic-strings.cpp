class Solution {
public:
    char cMap[127];
    unsigned int cSet[4];
    
    bool isIsomorphic(string s, string t) {
        if (s.length() != t.length()) return false;
        
        memset(cMap, 0, sizeof cMap);
        memset(cSet, 0, sizeof cSet);
        
        for (int i = 0; i < s.length(); ++i) {
            char a = s[i], b = t[i];
            if (!cMap[a]) {
                cMap[a] = b;
                int st = b/32;
                unsigned int offset = b%32;
                unsigned int v = cSet[st] & ((unsigned int)1<<offset);
                if (v) return false;
                cSet[st] |= ((unsigned int)1<<offset);
            }
            if (cMap[a] != b) return false;
        }
        
        return true;
    }
};