package main

func copyRandomList3(head *Node) *Node {
	if head == nil {
		return head
	}
	listmap := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		listmap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		listmap[curr].Next = listmap[curr.Next]
		listmap[curr].Random = listmap[curr.Random]
		curr = curr.Next
	}
	return listmap[head]
}
