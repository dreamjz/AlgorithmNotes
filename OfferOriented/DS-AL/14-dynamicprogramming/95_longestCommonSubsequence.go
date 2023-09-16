package ofdp

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// 缓存
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

// 迭代 SC: O(min(m,n)) 二维数组只保留两行
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m < n {
		return longestCommonSubsequence(text2, text1)
	}

	// 缓存
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 计算
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[(i+1)%2][j+1] = dp[i%2][j] + 1
			} else {
				dp[(i+1)%2][j+1] = max(dp[i%2][j+1], dp[(i+1)%2][j])
			}
		}
	}

	return dp[m%2][n]
}

// 迭代 一维数组
func longestCommonSubsequence3(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m < n {
		return longestCommonSubsequence(text2, text1)
	}

	// 缓存
	dp := make([]int, n+1)

	// 计算
	for i := range text1 {
		prev := dp[0] // prev is f(i-1, j-1)
		for j := range text2 {
			cur := 0 // cur is f(i,j)
			if text1[i] == text2[j] {
				cur = prev + 1 // f(i,j) = f(i-1,j-1) + 1
			} else {
				// f(i,j) = max(f(i,j-1), f(i-1, j))
				// dp[j] is f(i, j-1)
				// dp[j+1] is f(i-1, j)
				cur = max(dp[j], dp[j+1])
			}

			// prev is f(i-1, j)
			prev = dp[j+1]
			// d[j+1] is f(i, j)
			dp[j+1] = cur
		}
	}

	return dp[n]
}
