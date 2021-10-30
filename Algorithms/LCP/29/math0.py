class Solution:
    def orchestraLayout(self, num: int, xPos: int, yPos: int) -> int:
        x, y = yPos, xPos
        if x>=y and x<num-y:
            i = y*2
            st = i*(2*num-i)
            p = st + x-y
            return (p%9) + 1
        if x>=y and x>=num-y:
            i = (num-1 - x)*2
            ed = i*(2*num-i-2) + 2*num-2
            p = ed - (x-y)
            return (p%9) + 1
        if x<y and x>=num-1-y:
            i = 2*(num-y) - 1
            st = i*(2*num-i)
            p = st + y-1-x
            return (p%9) + 1
        if x<y and x<num-y:
            i = 2*x + 1
            ed = i*(2*num-i-2) + 2*num-2
            p = ed - (y - (x+1))
            return (p%9) + 1
        return -1