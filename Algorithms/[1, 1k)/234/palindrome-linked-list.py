# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        mod = 10**6 + 7
        left, right, mul = 0, 0, 1
        while head != None:
            left = (left*10 + head.val) % mod
            right = (right + head.val*mul) % mod
            mul = (mul*10) % mod
            head = head.next
        return left == right
