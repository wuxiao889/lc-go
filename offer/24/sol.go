package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		cn := curr.Next
		curr.Next = pre
		pre, curr = curr, cn
	}
	return pre
}
