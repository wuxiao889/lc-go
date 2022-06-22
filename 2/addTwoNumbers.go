package goleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	f := func(l *ListNode) int {
		if l == nil {
			return 0
		}
		return l.Val
	}
	var curr *ListNode
	var over int
	for l1 != nil || l2 != nil {
		v1 := f(l1)
		v2 := f(l2)
		val := (v1 + v2 + over) % 10
		if pre.Next == nil {
			pre.Next = &ListNode{val, nil}
			curr = pre.Next
		} else {
			curr.Next = &ListNode{val, nil}
			curr = curr.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		over = (v1 + v2 + over) / 10
	}
	if over == 1 {
		curr.Next = &ListNode{1, nil}
		curr = curr.Next
	}
	return pre.Next
}
