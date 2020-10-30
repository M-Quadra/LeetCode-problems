/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ptrA, ptrB := headA, headB

	cntA := 0
	for ; ptrA.Next != nil; ptrA = ptrA.Next {
		cntA++

	}

	cntB := 0
	for ; ptrB.Next != nil; ptrB = ptrB.Next {
		cntB++
	}

	if ptrA != ptrB {
		return nil
	}

	for ptrA = headA; cntA > cntB; cntA-- {
		ptrA = ptrA.Next
	}
	for ptrB = headB; cntB > cntA; cntB-- {
		ptrB = ptrB.Next
	}
	for ptrA != ptrB {
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}
	return ptrA
}