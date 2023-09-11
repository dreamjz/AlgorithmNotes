package orhash

import "strconv"

func findMinDifference(timePoints []string) int {
	totalMin := 24 * 60
	// 边界条件
	if len(timePoints) >= totalMin {
		return 0
	}

	// 创建哈希表
	minFlags := make([]int, totalMin)
	// 记录情况
	for _, t := range timePoints {
		hour, _ := strconv.Atoi(t[0:2])
		min, _ := strconv.Atoi(t[3:])
		mins := hour*60 + min
		// 出现重复的时间
		if minFlags[mins] == 1 {
			return 0
		}
		minFlags[mins] = 1
	}

	// 计算时间差
	prev := -1                   // 前一时间
	first := len(minFlags) - 1   // 最早
	last := -1                   // 最晚
	minDiff := len(minFlags) - 1 // 最小时间差
	for i := range minFlags {
		if minFlags[i] == 1 {
			if prev != -1 {
				minDiff = min(i-prev, minDiff)
			}
			prev = i
			first = min(i, first)
			last = max(i, last)
		}
	}
	// 最早最晚的时间差
	minDiff = min(first+len(minFlags)-last, minDiff)

	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
