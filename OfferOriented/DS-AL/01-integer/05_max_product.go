package soint

func maxProductBruteForce(strs []string) int {
	product := 0
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if compareStr(strs[i], strs[j]) {
				p := len(strs[i]) * len(strs[j])
				if product < p {
					product = p
				}
			}
		}
	}

	return product
}

func compareStr(a, b string) bool {
	// compare character one by one
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return false
			}
		}
	}
	return true
}

func maxProductHashTable(strs []string) int {
	// crate hash table
	ht := make([][]bool, len(strs))
	for i := range ht {
		ht[i] = make([]bool, 26)
	}
	for i := range strs {
		for _, s := range strs[i] {
			ht[i][s-'a'] = true
		}
	}

	// compare strings
	result := 0
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			k := 0
			// find common character
			for ; k < 26; k++ {
				if ht[i][k] && ht[j][k] {
					break
				}
			}
			// not found
			if k == 26 {
				p := len(strs[i]) * len(strs[j])
				if result < p {
					result = p
				}
			}
		}
	}

	return result
}

func maxProductBinary(strs []string) int {
	// create binary slice
	b := make([]int32, len(strs))
	for i := 0; i < len(strs); i++ {
		for _, c := range strs[i] {
			b[i] |= 1 << (c - 'a')
		}
	}

	// compare strings
	result := 0
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if b[i]&b[j] == 0 {
				p := len(strs[i]) * len(strs[j])
				if result < p {
					result = p
				}
			}
		}
	}

	return result
}
