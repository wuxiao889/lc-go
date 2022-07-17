package swapParis

import . "github.com/wuxiao889/goleetcode/util"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next.Next = swapPairs(head.Next.Next)
	node := head.Next
	head.Next, head.Next.Next = head.Next.Next, head
	return node
}
