package reverseList

import . "github.com/wuxiao889/goleetcode/util"

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		node := cur.Next
		cur.Next = pre
		pre = cur
		cur = node
	}
	return pre
}
