from typing import List

class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        strAry = list(S)
        optAry: List[str] = []
        def dfs(i):
            if i == len(strAry):
                optAry.append("".join(strAry))
                return
            
            v = strAry[i]
            if v.isdigit():
                dfs(i+1)
                return
            
            strAry[i] = v.lower()
            dfs(i+1)
            strAry[i] = v.upper()
            dfs(i+1)
        dfs(0)
        return optAry