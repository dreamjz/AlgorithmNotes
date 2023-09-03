package dsalarr

import "sort"

func threeSum(nums []int) [][]int {
	// qSortArr array in ascending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := make([][]int, 0)
	// threeSum
	for i := 0; i < len(nums); {
		result = twoSum07(nums, i, result)
		// skip identical element
		ie := nums[i]
		for i < len(nums) && nums[i] == ie {
			i++
		}

	}

	return result
}

func twoSum07(nums []int, i int, result [][]int) [][]int {
	// skip i
	left, right := i+1, len(nums)-1

	for left < right {
		sum := nums[i] + nums[left] + nums[right]
		if sum < 0 {
			left++
		}
		if sum > 0 {
			right--
		}
		if sum == 0 {
			result = append(result, []int{nums[i], nums[left], nums[right]})
			// skip identical nums
			le := nums[left]
			for left < right && nums[left] == le {
				left++
			}
		}
	}

	return result
}
