package ofsort

import (
	"math"
	"math/rand"
)

func SortArrayCountingSort(nums []int) []int {
	countingSort(nums)
	return nums
}

func countingSort(nums []int) {
	// get max/min number
	maxN, minN := math.MinInt, math.MaxInt
	for _, n := range nums {
		maxN = max(maxN, n)
		minN = min(minN, n)
	}

	// count numbers
	cnt := make([]int, maxN-minN+1)
	for _, n := range nums {
		cnt[n-minN]++
	}

	// rewrite array
	idx := 0
	for n := minN; n <= maxN; n++ {
		for cnt[n-minN] > 0 {
			nums[idx] = n
			idx++
			cnt[n-minN]--
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SortArrayQuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	// random pivot index
	pi := rand.Intn(right-left+1) + left
	swap(nums, pi, right) // move pivot to end

	// partition
	p := nums[right] // pivot value
	i, j := left-1, left
	for j < right {
		// exchange to left
		if nums[j] <= p {
			i++
			swap(nums, i, j)
		}

		j++
	}

	i++
	swap(nums, i, right)

	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func SortArrMergeSortRecursive(arr []int) []int {
	return MergeSortRecursive(arr)
}

func MergeSortRecursive(arr []int) []int {
	// only one element or empty
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	left := arr[0:mid] // left half
	right := arr[mid:] // right half

	// merge two halves
	return merge(MergeSortRecursive(left), MergeSortRecursive(right))
}

func merge(left, right []int) []int {
	// temporary slice
	maxLen := max(len(left), len(right))
	tmp := make([]int, 0, maxLen)

	// merge slices
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			tmp = append(tmp, left[0])
			left = left[1:]
		} else {
			tmp = append(tmp, right[0])
			right = right[1:]
		}
	}

	// add the rest
	if len(left) != 0 {
		tmp = append(tmp, left...)
	}
	if len(right) != 0 {
		tmp = append(tmp, right...)
	}

	return tmp
}

func SortArrMergeSortIterative(arr []int) []int {
	return MergeSortIterative(arr)
}

func MergeSortIterative(arr []int) []int {
	length := len(arr)
	res := make([]int, length)

	// merge adjacent intervals
	// double interval width every time
	for wid := 1; wid < length; wid *= 2 {
		// "start" is the first index of the first interval
		for start := 0; start < length; start += wid * 2 {
			mid := min(start+wid, length)   // end of first interval is mid-1
			end := min(start+wid*2, length) // end of next interval is end-1

			i := start // i is index of first interval
			j := mid   // j is index of next interval
			k := start // k is index of slice res
			// merge adjacent intervals
			for i < mid || j < end {
				if j == end || (i < mid && arr[i] <= arr[j]) {
					res[k] = arr[i]
					k++
					i++
				} else {
					res[k] = arr[j]
					k++
					j++
				}
			}
		}
		// swap two slice
		res, arr = arr, res
	}

	return arr
}
