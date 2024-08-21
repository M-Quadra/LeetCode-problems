class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> optAry;
        if (s.length() < 4 || 3*4 < s.length()) return optAry;
        
        auto siz = sizeof(bool)*s.length();
        bool *comAry = (bool *)malloc(siz);
        memset(comAry, 0, siz);
        memset(comAry + s.length() - 4, 1, sizeof(bool)*3);
        
        char str[20];
        
        do {
            auto tmpStr = s;
            for (int i = (int)s.length() - 1; ~i; --i) {
                if (!comAry[i]) continue;
                tmpStr.insert(i+1, ".");
            }
            
            memcpy(str, tmpStr.c_str(), sizeof(char)*tmpStr.length());
            str[tmpStr.length()] = 0;
            
            for (char *substr = strtok(str,":;."); substr != NULL; substr = strtok(NULL,":;.")) {
                
                auto num = stoi(substr);
                if (num > 255) goto ed;
                
                auto str = to_string(num);
                if (memcmp(str.c_str(), substr, sizeof(char)*strlen(substr))) {
                    goto ed;
                }
            }
            optAry.push_back(tmpStr);
        ed:;
        } while (next_permutation(comAry, comAry + s.length() - 1));
        
        return optAry;
    }
};