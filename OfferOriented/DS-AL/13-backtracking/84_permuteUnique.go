package ofbacktarck

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	helper84(nums, 0, &res)
	return res
}

func helper84(nums []int, idx int, res *[][]int) {
	if idx == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	} else {
		m := make(map[int]bool)
		for j := idx; j < len(nums); j++ {
			if !m[nums[j]] {
				m[nums[j]] = true
				swap(nums, j, idx)
				helper84(nums, idx+1, res)
				swap(nums, j, idx)
			}
		}
	}
}
