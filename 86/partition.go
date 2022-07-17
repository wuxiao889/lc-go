package partition

import . "github.com/wuxiao889/goleetcode/util"

func partition(head *ListNode, x int) *ListNode {
	pre1, pre2 := &ListNode{}, &ListNode{}
	c1, c2 := pre1, pre2
	cur := head
	for cur != nil {
		if cur.Val < x {
			c1.Next = cur
			c1 = c1.Next
			cur, cur.Next = cur.Next, nil
		} else {
			c2.Next = cur
			c2 = c2.Next
			cur, cur.Next = cur.Next, nil
		}
	}
	c1.Next = pre2.Next
	return pre1.Next
}
