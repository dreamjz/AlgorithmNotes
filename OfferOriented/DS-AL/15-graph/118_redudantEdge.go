package ofgraph

import "math"

func findRedundantConnection(edges [][]int) []int {
	// 获取所有节点
	maxVertex := math.MinInt
	for _, r := range edges {
		for _, n := range r {
			maxVertex = max(maxVertex, n)
		}
	}

	// 并查集
	fathers := make([]int, maxVertex+1)
	for i := 1; i < maxVertex+1; i++ {
		fathers[i] = i
	}

	for _, r := range edges {
		if !union(fathers, r[0], r[1]) {
			return []int{r[0], r[1]}
		}
	}
	return []int{}
}
