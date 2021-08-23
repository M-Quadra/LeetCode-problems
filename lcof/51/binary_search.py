from typing import List
import bisect
# https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/comments/364371

class Solution:
    def reversePairs(self, nums: List[int]) -> int:
        ary: List[int] = []
        opt = 0
        for v in nums:
            i = bisect.bisect_left(ary, -v)
            opt += i
            ary[i:i] = [-v]
        return opt
