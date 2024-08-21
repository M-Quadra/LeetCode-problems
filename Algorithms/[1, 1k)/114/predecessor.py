# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: TreeNode) -> None:
        now = root
        while now != None:
            if now.left != None:
                nxt: TreeNode = now.left
                while nxt.right != None: nxt = nxt.right
                nxt.right = now.right
                now.right = now.left
                now.left = None
            now = now.right