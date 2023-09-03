package dsalarr

func subarraySumBruteForce(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum == k {
				count++
			}
		}
	}

	return count
}

func subarraySumBruteForceOptimized(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

func subarraySumHashTable(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	count := 0

	m[0] = 1
	sum := 0
	for _, v := range nums {
		sum += v
		if val, ok := m[sum-k]; ok {
			count += val
		}
		m[sum] += 1

	}

	return count
}
