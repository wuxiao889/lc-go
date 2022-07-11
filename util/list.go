package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func Arr2List(nums []int) *ListNode {
	pre := &ListNode{}
	cur := pre
	for i := range nums {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return pre.Next
}

func ListToSlice(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		//fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	return nums
}
