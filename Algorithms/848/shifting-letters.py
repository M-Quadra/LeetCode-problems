from typing import List

class Solution:
    def shiftingLetters(self, s: str, shifts: List[int]) -> str:
        for i in range(len(shifts)-2, -1, -1):
            shifts[i] += shifts[i+1]
        
        opt = ''
        for i, v in enumerate(s.encode(encoding='ascii')):
            v -= 97
            v += shifts[i]
            v %= 26
            v += 97
            opt += chr(v)
        return opt