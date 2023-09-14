package ofbacktarck

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// sort
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := [][]int{}
	helper82(candidates, target, 0, &[]int{}, &res)
	return res
}

func helper82(nums []int, target, idx int, c *[]int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(*c))
		copy(tmp, *c)
		*res = append(*res, tmp)
	} else if idx < len(nums) && target > 0 {
		// 跳过相同数字
		helper82(nums, target, getNext(nums, idx), c, res)

		// 添加数字
		*c = append(*c, nums[idx])
		helper82(nums, target-nums[idx], idx+1, c, res)

		// 回溯
		*c = (*c)[:len(*c)-1]
	}
}

func getNext(nums []int, idx int) int {
	i := idx
	for i < len(nums) && nums[i] == nums[idx] {
		i++
	}
	return i
}
