package Array

func twoSum1(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		left, right := i+1, n-1
		for left <= right {
			mid := (left + right) / 2
			num := target - numbers[i]
			if num == numbers[mid] {
				return []int{i + 1, mid + 1}
			} else if num < numbers[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{-1, -1}
}
