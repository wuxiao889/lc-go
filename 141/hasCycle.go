package hasCycle

import . "github.com/wuxiao889/goleetcode/util"

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		//fast,fast.Next不为nil
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/*
快慢指针，有环一定会绕圈，绕圈快慢指针就会相遇
*/
