package main

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	for curr != nil {
		currnext := curr.Next
		curr.Next = &Node{Val: curr.Val}
		curr.Next.Next = currnext
		curr = currnext
	}
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}
	currhead := head.Next
	curr = head
	for curr != nil {
		curr2 := curr.Next
		curr.Next = curr.Next.Next
		if curr.Next != nil {
			curr2.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	return currhead
}
