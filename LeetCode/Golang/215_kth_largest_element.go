package golang

import "math/rand"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	t := n - k
	pi := partition(nums, 0, n-1)

	for pi != t {
		if pi < t {
			pi = partition(nums, pi+1, n-1)
		} else {
			pi = partition(nums, 0, pi-1)
		}
	}

	return nums[pi]
}

func partition(nums []int, left, right int) int {
	pi := rand.Intn(right-left+1) + left
	swap(nums, pi, right)

	p := nums[right]
	lo, hi := left, right
	for lo < hi {
		for lo < hi && nums[lo] <= p {
			lo++
		}
		for lo < hi && nums[hi] > p {
			hi--
		}
		if lo < hi {
			swap(nums, lo, hi)
		}
	}

	return lo - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
