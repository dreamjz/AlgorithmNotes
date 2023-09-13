package ofbs

func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := (left + right) / 2

		if mid <= x/mid {
			if mid+1 > x/(mid+1) {
				return mid
			}

			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}
