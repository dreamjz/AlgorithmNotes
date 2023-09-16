package ofdp

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	// 缓存
	dp := make([]int, n+1)
	// 初始值
	dp[0] = 1

	// 计算
	for _, cs := range s {
		// 新一行的计算
		prev := dp[0]
		for j, ct := range t {
			cur := 0
			if cs == ct {
				cur = prev + dp[j+1]
			} else {
				cur = dp[j+1]
			}

			// 缓存计算前的值, 下次计算还要用到
			// 缓存的上一行的值: dp[i][j+1]
			prev = dp[j+1]
			// 保存计算结果
			dp[j+1] = cur
		}
	}

	return dp[n]
}
