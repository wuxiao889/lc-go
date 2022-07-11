package longestCommonSubsequence

// dp[i][j] 定义为text1[0:i-1] text2[0:j-1]中最长公共子序列长度
//
// 1001*1001*8 字节
var dp [1001][1001]int

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
