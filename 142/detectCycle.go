package detectCycle

import . "github.com/wuxiao889/goleetcode/util"

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

/*
f s 快慢指针走的路程
a head到链表入口a个节点
b 环一圈节点

1.f=2s
2.f=s+nb
-> s=nb , f=2nb

head到入环要a+nb步，slow走了nb，再走a步
fast从head出发，和slow一起走，相遇时就时a步
*/
