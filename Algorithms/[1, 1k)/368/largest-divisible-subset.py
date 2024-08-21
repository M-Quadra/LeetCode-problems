from typing import List

class Solution:
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        nums.sort()
        idxAry = [0 for _ in range(len(nums))]
        lenAry = [i for i in range(len(nums), 0, -1)]
        nums.append(0)
        lenAry.append(0)

        optAry = []
        def dfs(idx, l) -> int:
            nonlocal idxAry, lenAry, optAry

            maxL = l
            for i in range(idx+1, len(nums)):
                if nums[i] == 0 or nums[i]%nums[idx] != 0:
                    if l <= len(optAry):
                        maxL = max(maxL, l)
                        continue

                    optAry = []
                    for i in range(0, l): optAry.append(nums[idxAry[i]])
                    continue
                elif l + lenAry[i] <= len(optAry):
                    maxL = max(maxL, l + lenAry[i])
                    continue
                    
                idxAry[l] = i
                maxL = max(maxL, dfs(i, l+1))
            return maxL
        
        for i in range(len(nums)-1)[::-1]:
            idxAry[0] = i
            lenAry[i] = dfs(i, 1)
        return optAry