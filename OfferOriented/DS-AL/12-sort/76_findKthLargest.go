package ofsort

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	t := n - k
	pivot := partition(nums, 0, n-1)

	for pivot != t {
		// in left part
		if pivot > t {
			pivot = partition(nums, 0, pivot-1)
		}
		// in right part
		if pivot < t {
			pivot = partition(nums, pivot+1, n-1)
		}
	}

	return nums[pivot]
}
