package ofbs

func peakIndexInMountainArray(arr []int) int {
	// 二分
	left, right := 1, len(arr)-2

	for left <= right {
		mid := (left + right) / 2

		// peek found
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		// peek in right half
		if arr[mid] > arr[mid-1] {
			left = mid + 1
		} else { // peek in left half
			right = mid - 1
		}
	}

	return -1
}
