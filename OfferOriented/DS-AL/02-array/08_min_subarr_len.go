package dsalarr

func minSubArrayLen(nums []int, k int) int {
	// two pointers
	left, right := 0, 0

	minLen := 0
	sum := 0
	for ; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= k {
			minLen = minLength(minLen, right-left+1)
			left++
			// next sub
			sum -= nums[left]
		}
	}

	return minLen
}

func minLength(m, l int) int {
	if m > 0 {
		if l < m {
			return l
		}
		return m
	}
	return l
}
