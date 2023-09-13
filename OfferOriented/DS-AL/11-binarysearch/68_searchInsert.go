package ofbs

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// 二分查找
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] >= target {
			// 边界: 插入最左边
			// 或 满足 nums[mid-1] < t <= nums[mid]
			if mid == 0 || nums[mid-1] < target {
				return mid
			}

			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 插入最右边
	return len(nums)
}
