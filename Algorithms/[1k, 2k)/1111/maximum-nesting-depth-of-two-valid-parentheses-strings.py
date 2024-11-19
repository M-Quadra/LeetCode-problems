from typing import List

class Solution:
    def maxDepthAfterSplit(self, seq: str) -> List[int]:
        deepA, deepB= 0, 0
        opt = [0 for _ in range(len(seq))]
        for i, v in enumerate(seq):
            if v == '(':
                if deepA <= deepB:
                    deepA += 1
                else:
                    opt[i] = 1
                    deepB += 1
            else:
                if deepA >= deepB:
                    deepA -= 1
                else:
                    opt[i] = 1
                    deepB -= 1            
        return opt
