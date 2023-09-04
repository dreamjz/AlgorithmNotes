package ordastring

func checkInclusion(s1, s2 string) bool {
	// 判断特殊情况
	if len(s2) < len(s1) {
		return false
	}

	// 哈希表, 字符: 次数
	cnt := make([]int, 26)

	// 扫描 s1, 次数值减一
	for _, c := range s1 {
		cnt[c-'a']--
	}

	// 扫描 s2, 移动子区间
	left := 0
	for right, c := range s2 {
		x := c - 'a'
		cnt[x]++
		// 次数大于0, 减小区间
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		// 此时 cnt 之和一定是不大于0的
		// 若区间长度和s1相同, 那么cnt之和一定为0
		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}
