class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        if s == '()':
            return 1

        r, c = 0, 1
        for i, v in enumerate(s[1:], 1):
            c += 1 if v == '(' else -1
            if c == 0:
                r = i + 1
                break
            
        if r == len(s):
            return 2 * self.scoreOfParentheses(s[1:-1])
        return self.scoreOfParentheses(s[:r]) + self.scoreOfParentheses(s[r:])