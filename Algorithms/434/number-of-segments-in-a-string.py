class Solution:
    def countSegments(self, s: str) -> int:
        return len([x for x in filter(None, s.split())])