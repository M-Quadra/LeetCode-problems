class Solution:
    def dfs(self, s: str, p: str) -> tuple[int, int]:
        if len(s) <= 0 or len(p) <= 0:
            while len(p) >= 2 and p[1] == "*":
                p = p[2:]
            return len(s), len(p)

        if len(p) == 1:
            if s[0] == p[0] or p[0] == ".":
                return len(s) - 1, len(p) - 1
            return len(s), len(p)
        
        if len(p) == 2 and p[1] == "*":
            if p[0] == ".":
                return 0, 0
            
            while len(s) > 0 and s[0] == p[0]:
                s = s[1:]
            return len(s), 0
        
        if p[1] != "*":
            if s[0] == p[0] or p[0] == ".":
                return self.dfs(s[1:], p[1:])
            return len(s), len(p)
        if p[-1] != "*":
            if s[-1] == p[-1] or p[-1] == ".":
                return self.dfs(s[:-1], p[:-1])
            return len(s), len(p)
        
        if p[0] != ".":
            nextP = p[2:]
            if self.dfs(s, nextP) == (0, 0):
                return 0, 0
            for i in range(len(s)):
                if s[i] != p[0]:
                    return self.dfs(s[i:], nextP)
                if self.dfs(s[i+1:], nextP) == (0, 0):
                    return 0, 0
            return self.dfs("", nextP)
        if p[-2] != ".":
            nextP = p[:-2]
            if self.dfs(s, nextP) == (0, 0):
                return 0, 0
            for i in range(len(s)-1, -1, -1):
                if s[i] != p[-2]:
                    return self.dfs(s[:i+1], nextP)
                if self.dfs(s[:i], nextP) == (0, 0):
                    return 0, 0
            return self.dfs("", nextP)
        
        if p[0] == ".":
            nextP = p[2:]
            for i in range(len(s)):
                if self.dfs(s[i:], nextP) == (0, 0):
                    return 0, 0
            return len(s), len(p)
        if p[-1] == ".":
            nextP = p[:-2]
            for i in range(len(s), -1, -1):
                if self.dfs(s[:i], nextP) == (0, 0):
                    return 0, 0
            return len(s), len(p)
        return len(s), len(p)

    def isMatch(self, s: str, p: str) -> bool:
        return self.dfs(s, p) == (0, 0)