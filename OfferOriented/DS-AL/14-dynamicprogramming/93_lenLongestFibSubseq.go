package ofdp

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)

	// 哈希表, 快速寻找目标值
	m := make(map[int]int)
	for i, a := range arr {
		m[a] = i
	}

	// 缓存
	// dp[i][j] 表示斐波那契{..., ai, aj}的长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	res := 2 // 长度至少为2, 后续才能构成斐波那契数列
	// i: [0, n-2]
	// j: [1, n-1]
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			ai, aj := arr[i], arr[j]
			k := -1
			// 找出满足 ak + ai = aj 的 k
			if v, ok := m[aj-ai]; ok {
				k = v
			}
			if k != -1 && k < i { // 找到 ak
				dp[i][j] = dp[k][i] + 1
			} else { // ak 不在 ai 之前, {ai, aj} 单独构成, 为后续准备
				dp[i][j] = 2
			}

			res = max(res, dp[i][j])
		}
	}

	if res == 2 {
		res = 0
	}
	return res
}
