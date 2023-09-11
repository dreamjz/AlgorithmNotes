package orhash

func isAnagramWithArray(s string, t string) bool {
	// 边界条件
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return false
	}

	// 仅包含小写字母, 使用数组
	cnt := make([]int, 26)

	for _, c := range s {
		cnt[c-'a']++
	}

	for _, c := range t {
		if cnt[c-'a'] == 0 {
			return false
		}
		cnt[c-'a']--
	}

	return true
}

func isAnagramWithHashTable(s string, t string) bool {
	// 边界
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return false
	}

	// 使用哈希表
	cnt := make(map[rune]int)

	for _, c := range s {
		cnt[c]++
	}
	for _, c := range t {
		if cnt[c] == 0 {
			return false
		}
		cnt[c]--
	}

	return true
}
