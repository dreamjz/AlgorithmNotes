package ofgraph

func findCircleNum(isConnected [][]int) int {
	// 访问标记
	visited := make([]bool, len(isConnected))

	res := 0
	for i := range isConnected {
		if !visited[i] {
			bfs(isConnected, visited, i)
			res++
		}
	}

	return res
}

func bfs(isConnected [][]int, visited []bool, i int) {
	// queue
	q := []int{i}
	visited[i] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for j, n := range isConnected[cur] {
			// 相连 && 未访问
			if n == 1 && !visited[j] {
				q = append(q, j)
				visited[j] = true
			}
		}
	}
}

func findCircleNum2(isConnected [][]int) int {
	// 初始化并查集
	// 初始共有 n 个子集
	n := len(isConnected)
	fathers := make([]int, n)
	for i := range fathers {
		fathers[i] = i
	}

	// 扫描邻接矩阵
	// 合并子集
	res := n // 初始子集数
	for i := range isConnected {
		for j := range isConnected[i] {
			// 如果相连 && 能够合并
			if isConnected[i][j] == 1 && union(fathers, i, j) {
				// 合并成功, 子集数减少
				res--
			}
		}
	}

	return res
}

func union(fathers []int, i, j int) bool {
	fatherOfI := findFather(fathers, i)
	fatherOfJ := findFather(fathers, j)
	// 合并
	// 不属于同一子集
	if fatherOfI != fatherOfJ {
		fathers[fatherOfI] = fatherOfJ
		return true
	}
	return false
}

func findFather(fathers []int, i int) int {
	if fathers[i] != i {
		// 寻找根节点
		// 路径压缩
		fathers[i] = findFather(fathers, fathers[i])
	}
	return fathers[i]
}
