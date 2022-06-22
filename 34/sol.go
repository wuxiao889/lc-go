package main

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 || nums[left-1] != target {
		return -1
	}
	return left - 1
}
