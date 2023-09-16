package ofdp

func coinChange(coins []int, amount int) int {
	// 缓存
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1 // 存放最小值
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
