package mergeBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cnt := 0
	var pre *ListNode
	cur := list1
	var startnode, endnode *ListNode
	for cnt < b+1 {
		if cnt == a {
			startnode = pre
		}
		if cnt == b {
			endnode = cur.Next
			cur.Next = nil
		}
		pre = cur
		cur = cur.Next
		cnt++
	}
	startnode.Next = list2
	cur = list2
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = endnode
	return list1
}
