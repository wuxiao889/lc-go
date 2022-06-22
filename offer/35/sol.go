package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	for curr != nil {
		currNode := &Node{Val: curr.Val}
		currNode.Next = curr.Next
		curr.Next = currNode
		curr = currNode.Next
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
		currNode := curr.Next
		curr.Next = currNode.Next
		curr = currNode.Next
		if curr != nil {
			currNode.Next = curr.Next
		}
	}
	return currhead
}
