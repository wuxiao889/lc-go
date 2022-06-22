package main

func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)>>2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
