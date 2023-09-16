package ofdp

func minFlipsMonoIncr(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// 缓存
	// dp[0] 表示翻转后,末尾为0的最优解
	// dp[1] 表示翻转后, 末尾为1的最优解
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// 最小问题
	if s[0] == '0' {
		dp[0][0] = 0
		dp[1][0] = 1
	} else { // s[0] == '1'
		dp[0][0] = 1
		dp[1][0] = 0
	}

	for i := 1; i < n; i++ {
		cur := i % 2
		prev := (i - 1) % 2
		if s[i] == '0' {
			dp[0][cur] = dp[0][prev]
			dp[1][cur] = min(dp[0][prev], dp[1][prev]) + 1
		} else { // s[i] == '1'
			dp[0][cur] = dp[0][prev] + 1
			dp[1][cur] = min(dp[0][prev], dp[1][prev])
		}
	}

	last := (n - 1) % 2
	return min(dp[0][last], dp[1][last])
}
