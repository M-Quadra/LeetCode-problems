/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init() { self.val = 0; self.next = nil; }
 *     public init(_ val: Int) { self.val = val; self.next = nil; }
 *     public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
 * }
 */
class Solution {
    
    func mergeKLists(_ lists: [ListNode?]) -> ListNode? {
        var arr: [ListNode] = []
        for var list in lists {
            while let n = list {
                arr.append(n)
                list = n.next
            }
        }
        arr.sort { $0.val < $1.val }
        for i in arr.indices where i > 0 {
            arr[i-1].next = arr[i]
        }
        return arr.first
    }
}