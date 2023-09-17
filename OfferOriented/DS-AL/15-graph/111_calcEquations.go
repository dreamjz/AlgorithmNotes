package ofgraph

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 邻接表
	graph := buildGraph(equations, values)

	res := make([]float64, len(queries))
	// 求解
	for i := range queries {
		from := queries[i][0]
		to := queries[i][1]

		_, fromOk := graph[from]
		_, toOk := graph[to]
		if !fromOk || !toOk { // 不在图中, 无解
			res[i] = -1
		} else {
			// 访问标识
			visited := make(map[string]bool, len(equations)*2)
			res[i] = dfs(graph, visited, from, to)
		}
	}

	return res
}

func dfs(graph map[string]map[string]float64, visited map[string]bool, from, to string) float64 {
	if from == to {
		return 1
	}

	visited[from] = true
	for next, val := range graph[from] {
		// 未访问
		if !visited[next] {
			nextVal := dfs(graph, visited, next, to)
			if nextVal > 0 { // 有结果
				return val * nextVal
			}
		}
	}
	// 回溯, 还原状态
	visited[from] = false
	// 此路径 无解
	return -1
}

func buildGraph(eqs [][]string, vals []float64) map[string]map[string]float64 {
	maxCap := len(eqs) * 2 // 可能的最大节点数
	graph := make(map[string]map[string]float64, maxCap)

	// len(eqs) == len(vals)
	for i := 0; i < len(vals); i++ {
		var1, var2 := eqs[i][0], eqs[i][1]
		val := vals[i]

		// var1 / var2
		if _, ok := graph[var1]; !ok {
			m := make(map[string]float64, maxCap)
			graph[var1] = m
		}
		graph[var1][var2] = val

		// var2 / var1
		if _, ok := graph[var2]; !ok {
			m := make(map[string]float64, maxCap)
			graph[var2] = m
		}
		graph[var2][var1] = 1 / val
	}

	return graph
}
