package ofdp

func uniquePaths(m int, n int) int {
	// 缓存
	dp := make([]int, n)
	// 初始值
	for i := range dp {
		dp[i] = 1
	}

	// 计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}
