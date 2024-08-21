from typing import List
import math

class Solution:
    def next_permutation(self, a: List[int], st: int, ed: int) -> bool:
        if ed <= st:
            return False
        for i in range(ed - 2, -1, -1):
            j = i + 1
            if a[i] < a[j]:
                for k in range(ed - 1, -1, -1):
                    if a[k] > a[i]:
                        a[i], a[k] = a[k], a[i]
                        a[j:ed] = reversed(a[j:ed])
                        return True
            if i == st:
                a[st:ed] = reversed(a[st:ed])
                return False

    def ck(self, a: int, b: int) -> bool:
        v = round(math.sqrt(a+b))
        return v*v == a+b

    def numSquarefulPerms(self, A: List[int]) -> int:
        cntDic = {int:int}
        ckDic = {(int,int): bool}
        for i, v1 in enumerate(A):
            for _, v2 in enumerate(A[i+1:]):
                ok = self.ck(v1, v2)
                if ok:
                    cntDic[v1] = cntDic.get(v1, 1) + 1
                    cntDic[v2] = cntDic.get(v2, 1) + 1
                ckDic[v1, v2] = ok
                ckDic[v2, v1] = ok
        ok = False
        for v in list(cntDic.values())[1:]:
            if v >= len(A):
                ok = True
                break
        if not ok:
            return 0

        A.sort()
        opt = 0
        while True:
            opt += 1
            for i, _ in enumerate(A[1:], 1):
                if not ckDic[A[i-1], A[i]]:
                    opt -= 1
                    break
            if not self.next_permutation(A, 0, len(A)):
                break
        return opt