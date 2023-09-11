package orhash

func isAlienSorted(words []string, order string) bool {
	// 哈希表记录字母顺序
	orderArr := make([]int, len(order))
	for i, c := range order {
		orderArr[c-'a'] = i
	}

	// 比较单词
	for i := 0; i < len(words)-1; i++ {
		if !isSorted(words[i], words[i+1], orderArr) {
			return false
		}
	}

	return true
}

func isSorted(s1, s2 string, orderArr []int) bool {
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return orderArr[s1[i]-'a'] < orderArr[s2[i]-'a']
		}
	}

	return i == len(s1)
}
