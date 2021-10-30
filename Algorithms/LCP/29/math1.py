class Solution:
    def orchestraLayout(self, num: int, xPos: int, yPos: int) -> int:
        a, b = (min(xPos, num-1-yPos), 1) if yPos >= xPos else (min(yPos, num-1-xPos), -1)
        return (4*num*a - 4*a*a - 2*a + b*(xPos+yPos) + (b>>1&1)*4*(num-a-1))%9 + 1