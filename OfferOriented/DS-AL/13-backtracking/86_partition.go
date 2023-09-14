package ofbacktarck

func partition(s string) [][]string {
	res := [][]string{}
	helper86(s, 0, &[]string{}, &res)
	return res
}

func helper86(s string, start int, sub *[]string, res *[][]string) {
	// 寻找完毕
	if start == len(s) {
		tmp := make([]string, len(*sub))
		copy(tmp, *sub)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			*sub = append(*sub, s[start:i+1])
			helper86(s, i+1, sub, res)
			// 回溯
			*sub = (*sub)[:len(*sub)-1]
		}
	}
}

func isPalindrome(s string, i, j int) bool {
	start, end := i, j
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}
	return true
}
