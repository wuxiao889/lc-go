package main

func reversePrint1(head *ListNode) []int {
	arr := []int{}
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	result := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}
