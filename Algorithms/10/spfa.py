from queue import Queue
from typing import Tuple

class Solution:
    dp = {(0, 0): True}
    q: Queue[Tuple[int, int]] = Queue()

    def addNext(self, next: Tuple[int, int]) -> None:
        if self.dp.get(next, False): return
        self.dp[next] = True
        self.q.put(next)

    def isMatch(self, s: str, p: str) -> bool:
        self.dp = {(0, 0): True}
        self.q = Queue(); self.q.put((0, 0))

        while not self.q.empty():
            i, j = self.q.get()
            if j == len(p):
                if i == len(s): break
                continue

            if i == len(s):
                next = (i, j+2)
                if j+1 < len(p) and p[j+1] == "*": self.addNext(next)
                continue

            if s[i] == p[j] or p[j] == ".":
                if j+1 < len(p) and p[j+1] == "*":
                    self.addNext((i+1, j))
                    self.addNext((i+1, j+2))
                    self.addNext((i, j+2))
                    continue
                self.addNext((i+1, j+1))
                continue

            if j+1 < len(p) and p[j+1] == "*":
                self.addNext((i, j+2))
                continue
        return self.dp.get((len(s), len(p)), False)