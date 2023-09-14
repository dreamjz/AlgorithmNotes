package ofsort

import "sort"

func merge74(intervals [][]int) [][]int {
	res := [][]int{}

	// sort in acsending order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// merge intervals
	for i := 0; i < len(intervals); {
		tmp := []int{intervals[i][0], intervals[i][1]}

		// merge
		j := i + 1
		for j < len(intervals) && tmp[1] >= intervals[j][0] {
			tmp[1] = max(tmp[1], intervals[j][1])
			j++
		}

		res = append(res, tmp)

		i = j
	}

	return res
}
