package ofdp

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)

	sum := 0 // 数组元素和
	for _, n := range nums {
		sum += n
	}

	// 不能拆分成两个背包
	// 或 所有数加起来都小于 target
	if (sum+target)%2 != 0 || sum < target {
		return 0
	}

	// 背包容量
	t := (sum + target) / 2

	// 缓存
	dp := make([]int, t+1)
	// 初始值(最小问题解)
	dp[0] = 1

	// 计算
	for i := 1; i < n+1; i++ {
		// 从右向左计算, 避免覆盖原数据
		for j := t; j >= nums[i-1]; j-- {
			dp[j] += dp[j-nums[i-1]]
		}
	}

	return dp[t]
}
