package dsalarr

func numSubarrayLessThanK(nums []int, k int) int {
	left, right := 0, 0
	count := 0

	prod := 1
	for ; right < len(nums); right++ {
		prod *= nums[right]
		for left <= right && prod >= k {
			prod /= nums[left]
			left++
		}

		if left <= right {
			count += right - left + 1
		}
	}

	return count
}
