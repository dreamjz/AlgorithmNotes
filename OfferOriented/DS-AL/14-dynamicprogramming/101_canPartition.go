package ofdp

func canPartition(nums []int) bool {
	n := len(nums)
	t := 0 // 背包容量
	for _, n := range nums {
		t += n
	}
	// 无法拆分
	if t%2 == 1 {
		return false
	}
	t = t / 2

	// 缓存
	dp := make([]bool, t+1)
	dp[0] = true

	// 计算
	for i := 1; i < n+1; i++ {
		// 从右到左
		for j := t; j >= 1; j-- {
			// 放入背包
			// dp[j] = dp[j]
			// 不放入
			if !dp[j] && j >= nums[i-1] {
				dp[j] = dp[j-nums[i-1]]
			}
		}
	}

	return dp[t]
}
