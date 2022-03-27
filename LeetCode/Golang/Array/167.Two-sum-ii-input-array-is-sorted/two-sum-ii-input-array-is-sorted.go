package Array

// twoSumIIBinarySearch 使用 二分查找
func twoSumIIBinarySearch(numbers []int, target int) []int {
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

// twoSumIIDoublePointer 使用双指针
func twoSumIIDoublePointer(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
