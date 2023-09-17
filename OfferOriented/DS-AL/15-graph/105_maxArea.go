package ofgraph

func maxAreaOfIslandBFS(grid [][]int) int {
	r, c := len(grid), len(grid[0])

	// 标识已访问
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	maxArea := 0 // 最大岛屿面积
	// 扫描矩阵
	for r := range grid {
		for c := range grid[r] {
			// 岛屿节点 且 未访问
			if grid[r][c] == 1 && !visited[r][c] {
				// 计算面积
				area := getAreaBFS(grid, visited, r, c)

				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}

func getAreaBFS(grid [][]int, visited [][]bool, i, j int) int {
	row, col := len(grid), len(grid[0])
	// 移动方向 上下左右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 队列
	queue := make([][]int, 0)
	// 初始节点入队
	queue = append(queue, []int{i, j})
	visited[i][j] = true

	area := 0 // 面积
	// BFS
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]
		area++

		// 相邻节点入队
		for _, dir := range dirs {
			r := cur[0] + dir[0]
			c := cur[1] + dir[1]
			// 合法坐标
			isValid := r >= 0 && r < row && c >= 0 && c < col
			// 是岛屿节点 且 未访问
			if isValid && grid[r][c] == 1 && !visited[r][c] {
				// 入队
				queue = append(queue, []int{r, c})
				// 访问标记
				visited[r][c] = true
			}
		}
	}

	return area
}

func maxAreaOfIslandDFSStack(grid [][]int) int {
	// 访问标记
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	maxArea := 0
	// 扫描矩阵
	for r := range grid {
		for c := range grid[r] {
			// 岛屿节点 && 未访问
			if grid[r][c] == 1 && !visited[r][c] {
				area := getAreaDFS(grid, visited, r, c)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}

func getAreaDFS(grid [][]int, visited [][]bool, i, j int) int {
	// Queue
	st := make([][]int, 0)
	// 初始节点入栈
	st = append(st, []int{i, j})
	visited[i][j] = true

	// 移动方向, 上下左右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := 0
	row, col := len(grid), len(grid[0])
	// DFS, 搜索节点
	for len(st) > 0 {
		// 当前节点出队
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		area++

		// 相邻节点入队
		for _, dir := range dirs {
			r := cur[0] + dir[0]
			c := cur[1] + dir[1]
			isValid := r >= 0 && r < row && c >= 0 && c < col
			// 合法坐标 && 岛屿节点 && 未访问
			if isValid && grid[r][c] == 1 && !visited[r][c] {
				// 入队
				st = append(st, []int{r, c})
				// 标记访问
				visited[r][c] = true
			}
		}
	}

	return area
}
func maxAreaOfIslandDFSRecursive(grid [][]int) int {
	// 访问标记
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	maxArea := 0
	// 扫描矩阵
	for r := range grid {
		for c := range grid[r] {
			// 岛屿节点 && 未访问
			if grid[r][c] == 1 && !visited[r][c] {
				area := getAreaDFSRecursive(grid, visited, r, c)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}

func getAreaDFSRecursive(grid [][]int, visited [][]bool, i, j int) int {
	area := 1
	visited[i][j] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		isValid := r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
		if isValid && grid[r][c] == 1 && !visited[r][c] {
			area += getAreaDFSRecursive(grid, visited, r, c)
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
