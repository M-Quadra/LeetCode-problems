from typing import List

class Solution:
    def decode(self, encoded: List[int], first: int) -> List[int]:
        opt = [first]
        for i, v in enumerate(encoded): opt.append(v^opt[i])
        return opt
