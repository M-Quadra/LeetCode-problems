/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        for (auto nowPtr = head; nowPtr != NULL; nowPtr = nowPtr->next) {
            if (nowPtr->val == 0x7FFFFF) return true;
            nowPtr->val = 0x7FFFFF;
        }
        return false;
    }
};