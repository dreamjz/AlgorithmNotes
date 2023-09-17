package ofgraph

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(int)
	dfs = func(src int) {
		// 添加到路径
		path = append(path, src)
		// 到达最后节点
		if src == len(graph)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		} else {
			// 搜索
			for _, n := range graph[src] {
				dfs(n)
			}
		}
		// 回溯
		path = path[:len(path)-1]
	}
	dfs(0)

	return res
}
