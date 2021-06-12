from typing import List

class Solution:
    def wiggleSort(self, nums: List[int]) -> None:
        for i in range(1, len(nums)):
            if i & 1 == 1:
                if nums[i-1] < nums[i]:
                    nums[i-1], nums[i] = nums[i], nums[i-1]
            elif nums[i-1] > nums[i]:
                nums[i-1], nums[i] = nums[i], nums[i-1]
