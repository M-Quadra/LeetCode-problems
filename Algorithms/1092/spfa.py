from typing import List
from queue import Queue,PriorityQueue

class Solution:
    def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
        opt = str1 + str2

        lenDic = {(0, 0): 0}

        q = Queue()
        q.put(("", 0, 0))
        while not q.empty():
            now, i1, i2 = q.get()

            if i1 >= len(str1) or i2 >= len(str2):
                nOpt = now + str1[i1:] + str2[i2:]
                if len(nOpt) < len(opt):
                    opt = nOpt
                    lenDic[(len(str1), len(str2))] = len(opt)
                continue
            
            if str1[i1] == str2[i2]:
                nNow = now + str1[i1]
                oLen = lenDic.get((i1+1, i2+1), len(opt))
                if len(nNow) < oLen:
                    lenDic[(i1+1, i2+1)] = len(nNow)
                    q.put((nNow, i1+1, i2+1))
            
            nNow = now + str1[i1]
            oLen = lenDic.get((i1+1, i2), len(opt))
            if len(nNow) < oLen:
                lenDic[(i1+1, i2)] = len(nNow)
                q.put((nNow, i1+1, i2))

            nNow = now + str2[i2]
            oLen = lenDic.get((i1, i2+1), len(opt))
            if len(nNow) < oLen:
                lenDic[(i1, i2+1)] = len(nNow)
                q.put((nNow, i1, i2+1))
        return opt