class Solution:
    def scoreOfParentheses(self, S):
        stack = [0]
        for v in S:
            if v == '(':
                stack.append(0)
                continue
            
            v = stack.pop()
            stack[-1] += max(2 * v, 1)
        return stack.pop()