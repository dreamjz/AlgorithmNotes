package ofgraph

func numSimilarGroups(strs []string) int {
	n := len(strs)
	// 边界
	if n == 1 {
		return 1
	}

	// 邻接表
	graph := make(map[string][]string, len(strs))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				if _, ok := graph[strs[i]]; !ok {
					graph[strs[i]] = []string{}
				}
				graph[strs[i]] = append(graph[strs[i]], strs[j])
			}
		}
	}

	// 访问标记
	visited := make(map[string]bool, n)
	res := 0
	for _, str := range strs {
		for !visited[str] {
			bfs117(graph, visited, str)
			res++
		}
	}

	return res
}

func bfs117(graph map[string][]string, visited map[string]bool, str string) {
	// Queue
	q := []string{str}
	visited[str] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, adj := range graph[cur] {
			if !visited[adj] {
				q = append(q, adj)
				visited[adj] = true
			}
		}
	}
}

func isSimilar(s1, s2 string) bool {
	difCnt := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			difCnt++
		}
	}

	return difCnt <= 2
}

func numSimilarGroups2(strs []string) int {
	n := len(strs)
	if n == 1 {
		return 1
	}
	// 并查集
	fathers := make([]int, n)
	for i := range fathers {
		fathers[i] = i
	}

	// 合并子集
	res := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) && union(fathers, i, j) {
				res--
			}
		}
	}

	return res
}
