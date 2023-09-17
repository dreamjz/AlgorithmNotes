package ofgraph

import "math"

func updateMatrix(mat [][]int) [][]int {
	dists := make([][]int, len(mat))
	for i := range dists {
		dists[i] = make([]int, len(mat[0]))
	}

	// Queue
	queue := [][]int{}
	// 初始化
	for r := range mat {
		for c := range mat[r] {
			if mat[r][c] == 0 {
				// 入队
				queue = append(queue, []int{r, c})
				// 最短距离 0
				dists[r][c] = 0
			} else { // 非0节点
				// 初始化为最大整数
				dists[r][c] = math.MaxInt
			}
		}
	}

	// 移动方向 上下左右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// BFS
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]

		dist := dists[cur[0]][cur[1]]
		// 相邻节点
		for _, dir := range dirs {
			r := cur[0] + dir[0]
			c := cur[1] + dir[1]
			// 合法坐标
			isValid := r >= 0 && r < len(mat) && c >= 0 && c < len(mat[0])
			if isValid {
				// 未访问
				if dists[r][c] > dist+1 {
					dists[r][c] = dist + 1
					// 入队
					queue = append(queue, []int{r, c})
				}
			}
		}
	}

	return dists
}
