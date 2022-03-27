package array

// bruteForce 暴力解法
func bruteForce(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hashTable 使用 哈希表
func hashTable(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if j, ok := hashTable[target-x]; ok {
			return []int{i, j}
		}
		hashTable[x] = i
	}
	return nil
}
