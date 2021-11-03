class Solution:
    def multiply(self, A: int, B: int) -> int:
        if B <= 0 : return 0
        if B == 1: return A
        return (self.multiply(A, B>>1)<<1) + (A if (B>>1)<<1 != B else 0)