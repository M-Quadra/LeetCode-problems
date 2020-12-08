from typing import List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def listOfDepth(self, tree: TreeNode) -> List[ListNode]:
        opt:ListNode = []
        que = [tree]

        for i in range(0x7FFFFF):
            head: ListNode = None
            list: ListNode = None
            for _ in range(1<<i):
                now = que[0]
                que = que[1:]
                if now == None:
                    que.append(None)
                    que.append(None)
                    continue

                que.append(now.left)
                que.append(now.right)
                if head == None:
                    head = ListNode(now.val)
                    list = head
                    continue

                list.next = ListNode(now.val)
                list = list.next
            if head == None:
                break

            opt.append(head)
        
        return opt