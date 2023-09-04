package ordastring

func findAnagrams(s, p string) []int {
	// 判断边界
	if len(s) < len(p) {
		return []int{}
	}

	result := make([]int, 0)
	// 哈希表, 字符:次数
	cnt := make([]int, 26)

	// 扫描 p
	for _, c := range p {
		cnt[c-'a']--
	}

	// 扫描 s, 移动子区间
	left := 0
	for right, c := range s {
		x := c - 'a'
		cnt[x]++
		// 右移左指针, 直到cnt[x]不大于0
		for cnt[x] > 0 {
			cnt[s[left]-'a']--
			left++
		}
		// 当子区间长度等于p长度时, 满足p及其变位词为s的子串
		if right-left+1 == len(p) {
			result = append(result, left)
		}
	}

	return result
}
