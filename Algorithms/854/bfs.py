from queue import Queue
from typing import Tuple

class Solution:
    def kSimilarity(self, s1: str, s2: str) -> int:
        dic = {s1: max(0, len(s1)-1)}
        s2Ary = list(s2)

        q:Queue[Tuple[str, int, int]] = Queue()
        q.put((s1, 0, 0))
        while not q.empty():
            s1, i, c = q.get_nowait()
            if len(s1) <= 0: continue

            if s1[0] == s2Ary[i]:
                if c >= dic.get(s1[1:], 0x7FFFFF): continue

                dic[s1[1:]] = c
                q.put((s1[1:], i+1, c))
                continue

            for j in range(len(s1)):
                if s1[j] != s2[i]: continue

                ary = list(s1)
                ary[j] = ary[0]
                nS1 = ''.join(ary[1:])
                if c+1 >= dic.get(nS1, 0x7FFFFF): continue

                dic[nS1] = c+1
                q.put((nS1, i+1, c+1))
        return dic['']