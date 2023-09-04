package ordastring

func countSubstrings(s string) int {
	count := 0
	// 边界
	if len(s) == 0 {
		return count
	}

	// 扫描 s
	for i := range s {
		// 奇数长回文
		count += countPalindrome(s, i, i)
		// 偶数长回文
		count += countPalindrome(s, i, i+1)
	}

	return count
}

func countPalindrome(s string, c1, c2 int) int {
	count := 0
	left, right := c1, c2
	// 从中心向两边扩展
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return count
		}
		count++
		left--
		right++
	}
	return count
}
