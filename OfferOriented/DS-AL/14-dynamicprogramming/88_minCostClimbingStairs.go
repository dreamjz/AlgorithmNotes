package ofdp

func minCostClimbingStairsRecursive(cost []int) int {
	length := len(cost)
	// 登顶的最小消耗
	return min(helperRecursive(cost, length-2), helperRecursive(cost, length-1))
}

// 状态转移方程
func helperRecursive(cost []int, idx int) int {
	// 递归出口
	if idx <= 1 {
		return cost[idx]
	}

	// 第idx的最小消耗
	return min(helperRecursive(cost, idx-2), helperRecursive(cost, idx-1)) + cost[idx]
}

func minCostClimbingStairsRecursiveWithBuf(cost []int) int {
	// 缓存
	dp := make([]int, len(cost))
	// 子问题
	helperRecursiveWithBuf(cost, dp, len(cost)-2)
	helperRecursiveWithBuf(cost, dp, len(cost)-1)

	return min(dp[len(cost)-2], dp[len(cost)-1])
}

func helperRecursiveWithBuf(cost, dp []int, idx int) {
	if idx <= 1 { // 0 or 1
		dp[idx] = cost[idx]
	} else if dp[idx] == 0 { // 还未求解
		// 求解子问题
		helperRecursiveWithBuf(cost, dp, idx-2)
		helperRecursiveWithBuf(cost, dp, idx-1)

		dp[idx] = min(dp[idx-2], dp[idx-1]) + cost[idx]
	}
}

func minCostClimbingStairsIterativeOn(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}

	return min(dp[n-2], dp[n-1])
}

func minCostClimbingStairsIterativeO1(cost []int) int {
	n := len(cost)
	// 缓存
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		a := (i - 2) % 2
		b := (i - 1) % 2
		dp[i%2] = min(dp[a], dp[b]) + cost[i]
	}

	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
