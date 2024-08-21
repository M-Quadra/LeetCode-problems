from typing import List
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head == None:
            return None

        valSet = {head.val}
        last = head
        now = head.next
        while now != None:
            if now.val in valSet:
                last.next = now.next
                now.next = None
                now = last.next
                continue
            
            valSet.add(now.val)
            last = now
            now = now.next
        return head