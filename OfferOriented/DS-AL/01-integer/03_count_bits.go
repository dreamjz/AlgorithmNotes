package soint

func countBitsOnk(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = countBitsOkOfOneNum(i)
	}
	return result
}

func countBitsOkOfOneNum(num int) int {
	count := 0
	for i := num; i != 0; {
		count++
		i = i & (i - 1)
	}
	return count
}

func countBitsOn(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}

	return result
}

func countBitsOnWithBitOperation(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}
