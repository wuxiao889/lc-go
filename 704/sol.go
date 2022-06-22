package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
