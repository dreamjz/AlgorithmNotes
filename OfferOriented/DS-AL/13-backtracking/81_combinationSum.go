package ofbacktarck

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	helper81(candidates, target, 0, &[]int{}, &res)
	return res
}

func helper81(nums []int, target, idx int, com *[]int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(*com))
		copy(tmp, *com)
		*res = append(*res, tmp)
	} else if idx < len(nums) && target > 0 {
		// 不添加数字
		helper81(nums, target, idx+1, com, res)

		// 添加数字
		*com = append(*com, nums[idx])
		// 数字可复用
		helper81(nums, target-nums[idx], idx, com, res)
		// 回溯
		*com = (*com)[:len(*com)-1]
	}
}
