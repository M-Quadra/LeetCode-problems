class Solution:
    def brokenCalc(self, X: int, Y: int) -> int:
        if Y <= X:
            return X - Y

        cnt = 0
        while True:
            if Y == X:
                break

            cnt += 1
            if (Y & 1) == 1:
                Y += 1
                continue

            if Y > X:
                Y >>= 1
            else:
                Y += 1
        return cnt