package ofbacktarck

func combine(n int, k int) [][]int {
	res := [][]int{}
	helper80(n, k, 1, &[]int{}, &res)
	return res
}

func helper80(n, k, idx int, com *[]int, res *[][]int) {
	if len(*com) == k {
		tmp := make([]int, len(*com))
		copy(tmp, *com)
		*res = append(*res, tmp)
	} else if idx <= n {
		// 不添加
		helper80(n, k, idx+1, com, res)

		// 添加一个数字
		*com = append(*com, idx)
		helper80(n, k, idx+1, com, res)
		// 回溯
		*com = (*com)[:len(*com)-1]
	}
}
