package maxSubArray

func maxSubArray(nums []int) int {
	//溢出
	//sum, res := -math.MaxInt, -math.MaxInt
	sum, res := -100000, -100000
	for i := range nums {
		if nums[i]+sum < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	return findMaxSubArray(nums, 0, len(nums)-1)
}

func findMaxSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	leftsum := findMaxSubArray(nums, left, mid)
	rightsum := findMaxSubArray(nums, mid+1, right)
	midsum := findMidMaxSubArray(nums, left, right)
	return max(leftsum, rightsum, midsum)
}

func findMidMaxSubArray(nums []int, left int, right int) int {
	leftMax := -100000
	sum := 0
	mid := left + (right-left)>>1
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftMax = max(leftMax, sum)
	}
	rightMax := -100000
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightMax = max(rightMax, sum)
	}
	return leftMax + rightMax
}

func max(data ...int) int {
	max := data[0]
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}
