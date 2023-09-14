package ofbacktarck

func permute(nums []int) [][]int {
	res := [][]int{}
	helper83(nums, 0, &res)
	return res
}

func helper83(nums []int, idx int, res *[][]int) {
	// 所有数字已选定
	if idx == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	} else {
		// 第i个位置的数字需要从 [i, n-1] 中选
		for j := idx; j < len(nums); j++ {
			swap(nums, idx, j)
			// 选定下一个
			helper83(nums, idx+1, res)
			// 回溯
			swap(nums, idx, j)
		}
	}
}

func swap(nums []int, i, j int) {
	if i != j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
