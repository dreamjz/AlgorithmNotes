package ofdp

func minimumTotal(triangle [][]int) int {
	// 三角形的底边长度
	n := len(triangle[len(triangle)-1])

	// 缓存
	dp := make([]int, n)

	// 计算
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else if i == j {
				dp[j] = dp[j-1] + triangle[i][j]
			} else if i > j {
				dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
			}
		}
	}

	res := dp[0]
	for _, v := range dp {
		res = min(res, v)
	}
	return res
}
