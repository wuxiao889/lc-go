package sortList

import (
	. "github.com/wuxiao889/goleetcode/util"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//无头结点的快慢指针，当只有两个节点时，会陷入死循环

	//var pre *ListNode
	//slow, fast := head, head
	//for fast != nil && fast.Next != nil {
	//	pre = slow
	//	fast = fast.Next.Next
	//	slow = slow.Next
	//}
	//head2 := slow
	//pre.Next = nil
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	head2 := slow.Next
	slow.Next = nil

	head = sortList(head)
	head2 = sortList(head2)

	return mergeList(head, head2)
}

func mergeList(head1, head2 *ListNode) *ListNode {
	//虚拟头节点
	dummy := &ListNode{}
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	//这里要注意不能用for
	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}

	node := dummy.Next
	dummy.Next = nil
	return node
}
