package ordastring

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	// 边界值
	if len(s) == 0 {
		return maxLen
	}

	// 哈希表, 字符:次数, ASCII共256个
	cnt := make([]int, 256)

	// 扫描 s, 移动子区间
	left := 0
	for right, c := range s {
		cnt[c]++
		// 包含重复字符则减小区间
		for hasGreaterThan1(cnt) {
			// 移除字符
			cnt[s[left]]--
			left++
		}
		// 更新最大长度
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func lengthOfLongestSubstringWithCountDup(s string) int {
	maxLen := 0

	// 边界值
	if len(s) == 0 {
		return maxLen
	}

	// 哈希表, 字符:次数
	cnt := make([]int, 256)

	// 扫描s, 移动子区间
	left := 0
	cntDup := 0 // 标记是否存在重复字符
	for right, c := range s {
		cnt[c]++
		if cnt[c] > 1 {
			cntDup++
		}
		// 存在重复字符,减小区间
		for cntDup > 0 {
			// 删除字符
			cnt[s[left]]--
			// 下个区间是否存在重复字符
			if cnt[s[left]] == 1 {
				cntDup--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func hasGreaterThan1(cnt []int) bool {
	for _, n := range cnt {
		if n > 1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
