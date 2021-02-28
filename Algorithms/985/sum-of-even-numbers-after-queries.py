from typing import List

class Solution:
    def sumEvenAfterQueries(self, A: List[int], queries: List[List[int]]) -> List[int]:
        optAry: List[int] = []
        lastSum = sum([x for x in filter(lambda x: x&1 == 0, A)])
        for v in queries:
            val, idx = v[0], v[1]
            oddA, oddV = A[idx]&1 == 1, val&1 == 1
            if oddA and oddV:
                lastSum += A[idx]+val
            elif not oddA and oddV:
                lastSum -= A[idx]
            elif not oddA and not oddV:
                lastSum += val
            optAry.append(lastSum)
            A[idx] += val
        return optAry