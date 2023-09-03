package dsalarr

func pivotIndex(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
		if (sum - nums[i]) == (total - sum) {
			return i
		}
	}

	return -1
}
