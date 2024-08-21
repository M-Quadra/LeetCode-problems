class Solution {
public:
    int facAry[10] = {1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880};
    
    vector<int> diffWaysToCompute(string input) {
        int mark = 100000;
        map<int, char> numIdxMap;
        map<int, string> strIdxMap;
        
        int numAry[1000], numCnt = 0;
        int strAry[1000], strCnt = 0;
        int oprAry[1000], oprCnt = 0;
        
        string str = "";
        int num = 0;
        char tmpC[2] = {0, 0};
        
        for (int i = 0; i < input.size(); ++i) {
            char c = input[i];
            if (c == '+' || c == '-' || c == '*') {
                int idx = (int)strIdxMap.size();
                strIdxMap[idx] = str;
                strAry[strCnt++] = idx;
                str = "";
                
                idx = (int)strIdxMap.size();
                strIdxMap[idx] = c;
                strAry[strCnt++] = idx;
                
                numAry[numCnt++] = num;
                num = 0;
                
                idx = (int)numIdxMap.size() + mark;
                numIdxMap[idx] = c;
                numAry[numCnt++] = idx;
                oprAry[oprCnt++] = idx;
            } else {
                num = num*10 + (c-'0');
                tmpC[0] = c;
                str.append(tmpC, 0, 1);
            }
        }
        if (num) {
            numAry[numCnt++] = num;
            
            int idx = (int)strIdxMap.size();
            strIdxMap[idx] = str;
            strAry[strCnt++] = idx;
        }
        
        set<string> optSet;
        vector<int> optAry;
        
        int subNumAry[1000], subNumCnt = numCnt;
        int subStrAry[1000], subStrCnt = strCnt;
        
        do {
            auto subStrIdxMap = strIdxMap;
            subNumCnt = numCnt;
            subStrCnt = strCnt;
            memcpy(subNumAry, numAry, sizeof(int)*numCnt);
            memcpy(subStrAry, strAry, sizeof(int)*strCnt);
            
            for (int i = 0; i < oprCnt; ++i) {
                int oprIdx = oprAry[i];
                int idx = 0;
                while (subNumAry[idx] != oprIdx) ++idx;
                char oprChar = numIdxMap[oprIdx];
                
                int numL = subNumAry[idx - 1];
                int numR = subNumAry[idx + 1];
                auto strL = subStrIdxMap[subStrAry[idx - 1]];
                auto strR = subStrIdxMap[subStrAry[idx + 1]];
                
                switch (oprChar) {
                    case '+':
                        num = numL + numR;
                        str = "(" + strL + "+" + strR + ")";
                        break;
                    case '-':
                        num = numL - numR;
                        str = "(" + strL + "-" + strR + ")";
                        break;
                    default:
                        num = numL * numR;
                        str = "(" + strL + "*" + strR + ")";
                        break;
                }
                
                memmove(subNumAry + idx - 1, subNumAry + idx + 1, sizeof(int)*(subNumCnt - idx - 1));
                memmove(subStrAry + idx - 1, subStrAry + idx + 1, sizeof(int)*(subStrCnt - idx - 1));
                
                subNumAry[idx - 1] = num;
                int mapIdx = (int)subStrIdxMap.size();
                subStrIdxMap[mapIdx] = str;
                subStrAry[idx - 1] = mapIdx;
                
                subNumCnt -= 2;
                subStrCnt -= 2;
            }
            
            if (optSet.find(str) != optSet.end()) continue;
            optSet.insert(str);
            
            optAry.push_back(num);
        } while (next_permutation(oprAry, oprAry + oprCnt));
        
        sort(optAry.begin(), optAry.end());
        return optAry;
    }
};
