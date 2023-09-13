package ofbs

func singleNonDuplicate(nums []int) int {
	// 二分查找分组
	left, right := 0, len(nums)/2
	for left <= right {
		mid := (left + right) / 2
		// 分组第一个元素
		idx := 2 * mid

		// 组内元素不同
		if idx < len(nums)-1 && nums[idx] != nums[idx+1] {
			// 第一个分组 或 第一个元素不同的分组
			if mid == 0 || nums[idx-2] == nums[idx-1] {
				return nums[idx]
			}

			// 目标在左
			right = mid - 1
		} else { // 分组元素相同, 目标在右
			left = mid + 1
		}
	}

	// 最后一组只有一个元素
	return nums[len(nums)-1]
}
