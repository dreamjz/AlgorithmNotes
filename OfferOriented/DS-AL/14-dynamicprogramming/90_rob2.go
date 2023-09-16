package ofdp

func rob2(nums []int) int {
	n := len(nums)
	// 边界
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 两个子问题
	// [0, n-2]
	res1 := robSub(nums[:n-1])
	// [1, n-1]
	res2 := robSub(nums[1:n])

	return max(res1, res2)
}

func robSub(sub []int) int {
	n := len(sub)
	// 缓存
	dp := make([]int, 2)
	// 最小子问题
	if n >= 1 {
		dp[0] = sub[0]
	}
	if n >= 2 {
		dp[1] = max(sub[0], sub[1])
	}

	for i := 2; i < n; i++ {
		a := (i - 2) % 2
		b := (i - 1) % 2
		dp[i%2] = max(dp[a]+sub[i], dp[b])
	}

	return dp[(n-1)%2]
}
