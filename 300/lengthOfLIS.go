package lengthOfLIS

//动态规划解法
func lengthOfLIS(nums []int) int {
	//dp定义为数组中下标为i的元素最长公共子序列长度
	dp := make([]int, len(nums))
	//初始化为1，每个元素一定是一个以自己为开始的上升子序列
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	var max int
	for i := range dp {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
