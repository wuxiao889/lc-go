package getIntersectionNode

import . "github.com/wuxiao889/goleetcode/util"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	c1, c2 := headA, headB
	for c1 != nil || c2 != nil {
		if c1 == c2 {
			break
		}
		if c1 == nil {
			c1 = headB
		} else {
			c1 = c1.Next
		}
		if c2 == nil {
			c2 = headA
		} else {
			c2 = c2.Next
		}
	}
	return c1
}
