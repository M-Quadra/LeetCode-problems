# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    sum, half, mul = 0, 0, 0

    def sumVal(self, now: TreeNode) -> int:
            if now == None:
                return 0
            
            sumL = self.sumVal(now.left)
            sumR = self.sumVal(now.right)
            now.val += sumL + sumR
            return now.val
    
    def dfs(self, now: TreeNode):
            if now == None:
                return
            
            nMul = self.sum - now.val
            if abs(nMul - self.half) < abs(self.mul - self.half):
                self.mul = nMul
            if now.val < self.half:
                return
            
            self.dfs(now.left)
            self.dfs(now.right)
        
    def maxProduct(self, root: TreeNode) -> int:
        self.sum = self.sumVal(root)
        self.half = self.sum >> 1 | 1
        self.mul = 0
        self.dfs(root.left)
        self.dfs(root.right)
        return self.mul*(self.sum - self.mul)  % (10**9 + 7)
