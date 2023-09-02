package dsalarr

func twoSumBruteForce(ints []int, n int) []int {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i]+ints[j] == n {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumBinarySearch(ints []int, n int) []int {
	for i := range ints {
		t := n - ints[i]
		j := binarySearch(ints, t)
		if j != -1 && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func binarySearch(ints []int, t int) int {
	left := 0
	right := len(ints) - 1

	for left <= right {
		mid := (left + right) / 2

		midVal := ints[mid]
		if t == midVal {
			return mid
		}
		if t > midVal {
			left++
		}
		if t < midVal {
			right--
		}
	}

	return -1
}

func twoSumHashTable(ints []int, n int) []int {
	m := make(map[int]int, len(ints))
	for i, v := range ints {
		m[v] = i
	}

	for i, v := range ints {
		if val, isPresent := m[n-v]; isPresent {
			if i != val {
				return []int{i, val}
			}
		}
	}

	return []int{}
}

func twoSumTwoPointer(ints []int, n int) []int {
	left := 0
	right := len(ints) - 1

	for left < right {
		sum := ints[left] + ints[right]
		if sum == n {
			return []int{left, right}
		}
		if sum > n {
			right--
		}
		if sum < n {
			left++
		}
	}

	return []int{}
}
