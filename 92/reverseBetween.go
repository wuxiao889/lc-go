package reverseBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//left==1时，等价于反转前right个
	if left == 1 {
		return reverseN(head, right)
	}
	//递归
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var rear *ListNode

//递归版
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		//保存后驱结点
		rear = head.Next
		return head
	}
	//新链表头
	node := reverseN(head.Next, n-1)
	//反转head.Next
	head.Next.Next = head
	//反转head,指向后驱
	head.Next = rear
	return node
}
