from queue import LifoQueue

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: TreeNode) -> None:
        now, s = root, LifoQueue[TreeNode]()
        while now != None:
            if now.left != None:
                if now.right != None:
                    s.put(now.right)
                now.right = now.left
                now.left = None
            elif now.right == None and not s.empty():
                now.right = s.get()
            now = now.right