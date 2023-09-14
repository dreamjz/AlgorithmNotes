package ofbacktarck

func subsets(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	helper(nums, 0, &[]int{}, &res)

	return res
}

func helper(nums []int, idx int, subset *[]int, res *[][]int) {
	if idx == len(nums) {
		tmp := make([]int, len(*subset))
		copy(tmp, *subset)
		*res = append(*res, tmp)
	} else {
		// 不添加元素
		helper(nums, idx+1, subset, res)

		// 添加元素
		*subset = append(*subset, nums[idx])
		helper(nums, idx+1, subset, res)
		// 回溯
		*subset = (*subset)[:len(*subset)-1]
	}
}
