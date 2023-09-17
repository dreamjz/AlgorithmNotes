package ofgraph

func isBipartiteDFS(graph [][]int) bool {
	// 着色标记
	colors := make([]int, len(graph))
	for i := range colors {
		colors[i] = -1
	}

	for i := range graph {
		if colors[i] == -1 {
			if !setColorDFS(graph, colors, i, 0) {
				return false
			}
		}
	}

	return true
}

func setColorDFS(graph [][]int, colors []int, i, c int) bool {
	// Queue
	st := make([]int, 0)
	// 初始节点入队
	st = append(st, i)
	// 标记
	colors[i] = c

	// BFS
	for len(st) > 0 {
		// 出栈
		l := len(st)
		cur := st[l-1]
		st = st[:l-1]

		// 相邻节点入栈
		for _, adj := range graph[cur] {
			// 已着色但颜色相同, 不是二分图
			if colors[adj] != -1 {
				if colors[adj] == colors[cur] {
					return false
				}
			} else { // 未着色
				// 入栈
				st = append(st, adj)
				// 添加不同色
				colors[adj] = 1 - colors[cur]
			}
		}
	}

	return true
}

func isBipartiteBFS(graph [][]int) bool {
	// 着色标记
	colors := make([]int, len(graph))
	for i := range colors {
		colors[i] = -1
	}

	for i := range graph {
		if colors[i] == -1 {
			if !setColorBFS(graph, colors, i, 0) {
				return false
			}
		}
	}

	return true
}

func setColorBFS(graph [][]int, colors []int, i, c int) bool {
	// Queue
	queue := make([]int, 0)
	// 初始节点入队
	queue = append(queue, i)
	// 标记
	colors[i] = c

	// BFS
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]

		// 相邻节点入队
		for _, adj := range graph[cur] {
			// 已着色但颜色相同, 不是二分图
			if colors[adj] != -1 {
				if colors[adj] == colors[cur] {
					return false
				}
			} else { // 未着色
				// 入队
				queue = append(queue, adj)
				// 添加不同色
				colors[adj] = 1 - colors[cur]
			}
		}
	}

	return true
}
