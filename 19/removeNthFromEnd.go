package removeNthFromEnd

import . "github.com/wuxiao889/goleetcode/util"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{Next: head}
	fast, slow := pre, pre
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
