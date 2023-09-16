package ofdp

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	// 边界
	if m+n != len(s3) {
		return false
	}

	// 缓存长度取较小值
	if m < n {
		return isInterleave(s2, s1, s3)
	}

	// 缓存
	dp := make([]bool, n+1)
	dp[0] = true

	// 最小问题, 初始值
	for i := 1; i < n+1; i++ {
		dp[i] = s2[:i] == s3[:i]
	}

	// 计算
	for i, c1 := range s1 {
		// 对应二维数组的 dp[i][0]
		// 相当于二维数组计算到新的一行
		dp[0] = dp[0] && (c1 == rune(s3[i]))
		for j, c2 := range s2 {
			c3 := rune(s3[i+j+1])
			// 从左到右
			// dp[j+1]: dp[i][j+1]
			// dp[j]: dp[i+1][j]
			dp[j+1] = (c1 == c3 && dp[j+1]) || (c2 == c3 && dp[j])
		}
	}

	return dp[n]
}
