from typing import List

class Solution:
    tre: List[int] = []
    def lowBit(self, i: int) -> int:
        return -i&i
    def add(self, i: int, v: int):
        while i < len(self.tre):
            self.tre[i] += v
            i += self.lowBit(i)
    def sum(self, i: int) -> int:
        opt = 0
        while i:
            opt += self.tre[i]
            i -= self.lowBit(i)
        return opt

    def reversePairs(self, nums: List[int]) -> int:
        uniAry = list(set(nums))
        uniAry.sort()
        self.tre = [0 for _ in range(len(uniAry)+1)]
        
        uniDic: dict[int, int] = {}
        cntDic: dict[int, int] = {}
        for i, v in enumerate(uniAry):
            uniDic[v] = i + 1
            cntDic[i + 1] = 0
        
        for i, v in enumerate(nums):
            v = uniDic[v]
            nums[i] = v
            cntDic[v] += 1
        for k, v in cntDic.items():
            self.add(k, v)
        
        opt = 0
        for v in nums:
            opt += self.sum(v-1)
            self.add(v, -1)
        return opt
