package ofdp

func robRecursiveBuf(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 缓存
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	// 子问题
	helperRecursiveBuf(nums, dp, n-1)

	// 最优解
	return dp[n-1]
}

// 状态转移方程
func helperRecursiveBuf(nums, dp []int, idx int) {
	// 最小问题
	if idx == 0 {
		dp[idx] = nums[0]
	} else if idx == 1 {
		dp[idx] = max(nums[0], nums[1])
	} else if dp[idx] == -1 { // 未计算过
		// 子问题
		helperRecursiveBuf(nums, dp, idx-2)
		helperRecursiveBuf(nums, dp, idx-1)

		// 当前问题最优解
		dp[idx] = max(dp[idx-2]+nums[idx], dp[idx-1])
	}
}

func robIterativeOn(nums []int) int {
	n := len(nums)
	// 边界
	if n == 0 {
		return 0
	}

	// 缓存
	dp := make([]int, n)
	// 最小问题
	if n >= 1 {
		dp[0] = nums[0]
	}
	if n >= 2 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}

func rob(nums []int) int {
	n := len(nums)
	// 边界
	if n == 0 {
		return 0
	}

	// 缓存
	dp := make([]int, 2)
	// 最小子问题
	if n >= 1 {
		dp[0] = nums[0]
	}
	if n >= 2 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < n; i++ {
		a := (i - 2) % 2
		b := (i - 1) % 2
		dp[i%2] = max(dp[a]+nums[i], dp[b])
	}

	return dp[(n-1)%2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
