class Solution:
    def kSimilarity(self, s1: str, s2: str) -> int:
        dic = {s1: max(0, len(s1)-1), s2: max(0, len(s1)-1)}
        s1Ary, s2Ary = list(s1), list(s2)
        def dfs(st: int, cnt: int) -> int:
            if st >= len(s1Ary): return cnt
            while st < len(s1Ary) and s1Ary[st] == s2Ary[st]:
                st += 1
            
            opt = cnt + max(0, len(s1)-st-1)
            for i in range(st+1, len(s1Ary)):
                if s1Ary[i] == s2Ary[st]:
                    s1Ary[st], s1Ary[i] = s1Ary[i], s1Ary[st]
                    nextS = ''.join(s1Ary)
                    oldCnt = dic.get(nextS, 0x7FFFFF)
                    if cnt < oldCnt:
                        dic[nextS] = cnt
                        opt = min(opt, dfs(st+1, cnt+1))
                    s1Ary[st], s1Ary[i] = s1Ary[i], s1Ary[st]
            return opt
        return dfs(0, 0)