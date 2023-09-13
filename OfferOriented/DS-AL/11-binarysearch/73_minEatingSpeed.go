package ofbs

import "math"

func minEatingSpeed(piles []int, h int) int {
	// find max pile
	maxPile := math.MinInt
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	// binary search
	left, right := 1, maxPile
	for left <= right {
		mid := (left + right) / 2

		// 能吃完
		if getTime(piles, mid) <= h {
			// 速度 1 或 能吃完的最小速度
			if mid == 1 || getTime(piles, mid-1) > h {
				return mid
			}

			// 减小速度
			right = mid - 1
		} else {
			// 加大速度
			left = mid + 1
		}
	}

	return -1
}

func getTime(piles []int, speed int) int {
	t := 0
	for _, p := range piles {
		if p%speed == 0 {
			t += p / speed
		} else {
			t += p/speed + 1
		}
	}
	return t
}
