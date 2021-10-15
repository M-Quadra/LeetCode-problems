from typing import List

class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        nums.sort()
        sumAry = nums.copy()
        for i in range(1, len(sumAry)):
            sumAry[i] += sumAry[i-1]
        nums.reverse()
        sumAry.reverse()

        dic = {0: 1}
        for i, n in enumerate(nums):
            nowDic = dic.copy()
            dic.clear()
            for k, v in nowDic.items():
                minN, maxN = k-sumAry[i], k+sumAry[i]
                if target < minN or maxN < target:
                    continue

                dic[k-n] = dic.get(k-n, 0) + v
                dic[k+n] = dic.get(k+n, 0) + v
        return dic.get(target, 0)