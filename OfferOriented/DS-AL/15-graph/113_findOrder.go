package ofgraph

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 构建有向图
	graph := make(map[int][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		s := make([]int, 0)
		graph[i] = s
	}
	// 邻接表
	inDegrees := make([]int, numCourses) // 节点入度
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		inDegrees[pre[0]]++
	}

	// 广度优先搜索入度为0的节点
	queue := make([]int, 0)
	for i, d := range inDegrees {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	// 拓扑排序序列
	res := make([]int, 0, numCourses)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 加入拓扑排序序列
		res = append(res, cur)
		// 删除以 cur 为起点的边
		for _, adj := range graph[cur] {
			inDegrees[adj]--
			// 添加新的入度为 0 的节点
			if inDegrees[adj] == 0 {
				queue = append(queue, adj)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}

	return []int{}
}
