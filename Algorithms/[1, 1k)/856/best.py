class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        opt, d = 0, 0
        for i, v in enumerate(s):
            d += 1 if v == '(' else -1
            if v == ')' and s[i-1:i+1] == '()':
                opt += 1 << d
        return opt