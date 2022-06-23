package main

func findRepeatNumber(nums []int) int {
	arr := make(map[int]int)
	for _, v := range nums {
		if arr[v] == 1 {
			return v
		} else {
			arr[v] = 1
		}
	}
	return 0
}
