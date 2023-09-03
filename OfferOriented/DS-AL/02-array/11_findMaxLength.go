package dsalarr

func findMaxLength(nums []int) int {
	maxLen := 0

	m := make(map[int]int, len(nums))
	m[0] = -1
	sum := 0
	for i, n := range nums {
		if n == 0 {
			n = -1
		}
		sum += n
		if j, ok := m[sum]; ok {
			maxLen = max(maxLen, i-j)
		} else {
			// 第一次和为 sum 的下标
			m[sum] = i
		}

	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
