package change

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i < len(coins)+1; i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i-1]]
		}
	}
	return dp[amount]
}

func change1(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	//dp[i][j] 若只使用 coins 中的前 i 个（i 从 1 开始计数）硬币的面值，若想凑出金额 j，有 dp[i][j] 种凑法。
	for i := 1; i < len(coins)+1; i++ {
		for j := 0; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}
