package ofbs

import "math/rand"

type Solution struct {
	sums  []int // 前i项权重和
	total int   // 权重和
}

func Constructor(w []int) Solution {
	sums := make([]int, 0, len(w))
	total := 0
	for _, n := range w {
		total += n
		sums = append(sums, total)
	}

	return Solution{sums: sums, total: total}
}

func (s *Solution) PickIndex() int {
	p := rand.Intn(s.total)

	// 二分查找第一个大于 p 的
	left, right := 0, len(s.sums)-1
	for left <= right {
		mid := (left + right) / 2

		if s.sums[mid] > p {
			// 首个元素 或 第一个大于 p
			if mid == 0 || s.sums[mid-1] <= p {
				return mid
			}

			// 左
			right = mid - 1
		} else {
			// 右
			left = mid + 1
		}
	}

	return -1
}
