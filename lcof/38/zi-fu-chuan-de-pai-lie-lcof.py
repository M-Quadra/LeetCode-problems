from typing import List

class Solution:
    def next_permutation(self, a: List[str], st: int, ed: int) -> bool:
        if ed <= st:
            return False
        for i in range(ed - 2, -1, -1):
            j = i + 1
            if a[i] < a[j]:
                for k in range(ed - 1, -1, -1):
                    if a[k] > a[i]:
                        a[i], a[k] = a[k], a[i]
                        a[j:ed] = a[j:ed][::-1]
                        return True
            if i == st:
                a[st:ed] = a[st:ed][::-1]
                return False
        return False

    def permutation(self, s: str) -> List[str]:
        ary = sorted(s)
        
        opt = [''.join(ary)]
        while self.next_permutation(ary, 0, len(s)):
            opt.append(''.join(ary))
        return opt