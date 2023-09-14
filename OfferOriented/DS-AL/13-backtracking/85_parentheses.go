package ofbacktarck

func generateParenthesis(n int) []string {
	res := []string{}
	helper85(n, n, "", &res)
	return res
}

func helper85(left, right int, par string, res *[]string) {
	// 生成完毕
	if left == 0 && right == 0 {
		*res = append(*res, par)
		return
	}

	// 左括号
	if left > 0 {
		helper85(left-1, right, par+"(", res)
	}

	// 右括号
	if left < right {
		helper85(left, right-1, par+")", res)
	}
}
