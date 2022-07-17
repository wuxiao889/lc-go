package deleteDuplicates

import (
	. "github.com/wuxiao889/goleetcode/util"
)

func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	cur := pre
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next, cur.Next.Next = cur.Next.Next, nil
			}
		} else {
			cur = cur.Next
		}
	}
	return pre.Next
	//node := &ListNode{math.MaxInt, head}
	//pre := node
	//cur := head
	//dup := false
	//for cur != nil && cur.Next != nil {
	//	if cur.Val == cur.Next.Val {
	//		dup = true
	//		cur.Next, cur.Next.Next = cur.Next.Next, nil
	//	} else {
	//		if dup == true {
	//			pre.Next = cur.Next
	//			cur, cur.Next = cur.Next, nil
	//			dup = false
	//		} else {
	//			pre, cur = pre.Next, cur.Next
	//		}
	//	}
	//}
	//if dup == true {
	//	pre.Next = cur.Next
	//	cur.Next = nil
	//}
	//return node.Next
}
