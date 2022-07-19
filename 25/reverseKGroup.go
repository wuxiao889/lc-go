package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代版
func reverseKGroup(head *ListNode, k int) *ListNode {
	var cnt int
	//尾节点
	var tail *ListNode
	cur := head
	//计数，检查当前链表有没有k个节点
	//有k个，tail指向第k,cur指向k+1
	for cur != nil && cnt < k {
		tail = cur
		cur = cur.Next
		cnt++
	}
	//没有k个不反转
	if cnt < k {
		return head
	}
	//递归
	tail.Next = reverseKGroup(tail.Next, k)
	//对前k个反转
	head = reverseN(head, k)
	return head
}

//迭代版反转
func reverseN(head *ListNode, k int) *ListNode {
	cnt := 0
	var dummy *ListNode
	pre, cur := dummy, head
	for cnt < k {
		//保存下一个节点
		next := cur.Next
		//反转
		cur.Next = pre
		//后移
		pre = cur
		cur = next
		cnt++
	}
	//cnt==k，cur指向第k+1,pre第k是新链表的头
	head.Next = cur
	return pre
}
