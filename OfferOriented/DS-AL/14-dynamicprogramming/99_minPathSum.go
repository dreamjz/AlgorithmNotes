package ofdp

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 缓存
	dp := make([]int, n)
	// 初始值, 最小问题解
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	// 计算
	for i := 1; i < m; i++ {
		// 新的一行
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[n-1]
}
