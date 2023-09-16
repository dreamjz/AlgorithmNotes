package ofdp

func minCost(costs [][]int) int {
	n := len(costs)
	// 边界
	if n == 0 {
		return 0
	}

	// 缓存
	dp := make([][]int, 3)
	// 最小问题
	for i := range dp {
		tmp := make([]int, 2)
		dp[i] = append(tmp, costs[0][i])
	}

	for i := 1; i < n; i++ {
		// 三种颜色
		for j := 0; j < 3; j++ {
			prev1 := dp[(j+2)%3][(i-1)%2]
			prev2 := dp[(j+1)%3][(i-1)%2]

			dp[j][i%2] = min(prev1, prev2) + costs[i][j]
		}
	}

	// 三种方案中最小的
	return min(dp[0][(n-1)%2], min(dp[1][(n-1)%2], dp[2][(n-1)%2]))
}
