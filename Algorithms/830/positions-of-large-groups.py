from typing import List

class Solution:
    def largeGroupPositions(self, s: str) -> List[List[int]]:
        opt: List[List[int]] = []
        s += '-'
        st = -1
        for i, v in enumerate(s[:-1]):
            if v == s[i+1]:
                if st == -1:
                    st = i
            elif st != -1:
                ed = i
                if ed - st > 1:
                    opt.append([st, ed])
                st = -1
        return opt