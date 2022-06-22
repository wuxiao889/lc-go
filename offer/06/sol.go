package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := []int{}
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
