class Solution:
    def calculate(self, s: str) -> int:
        x, y = 1, 0
        for _, v in enumerate(s):
            if v == 'A':
                x = 2*x + y
            else:
                y = 2*y + x
        return x + y