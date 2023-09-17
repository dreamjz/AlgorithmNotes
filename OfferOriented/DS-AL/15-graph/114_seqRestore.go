package ofgraph

import "reflect"

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	// 构建有向图
	inDegrees := make(map[int]int, len(nums)) // 入度
	graph := make(map[int][]int, len(nums))
	for _, seq := range sequences {
		if len(seq) <= 1 {
			if _, ok := graph[seq[0]]; !ok {
				s := make([]int, 0, len(nums))
				graph[seq[0]] = s
			}
			if _, ok := inDegrees[seq[0]]; !ok {
				inDegrees[seq[0]] = 0
			}
		}
		for i := 0; i < len(seq)-1; i++ {
			if _, ok := graph[seq[i]]; !ok {
				s := make([]int, 0, len(nums))
				graph[seq[i]] = s
			}
			graph[seq[i]] = append(graph[seq[i]], seq[i+1])
			// 入度增加
			if _, ok := inDegrees[seq[i]]; !ok {
				inDegrees[seq[i]] = 0
			}
			inDegrees[seq[i+1]]++
		}
	}

	// 入度为0的节点
	queue := make([]int, 0)
	for k, v := range inDegrees {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	// 拓扑排序序列
	res := make([]int, 0, len(nums))
	// 寻找唯一拓扑排序序列
	for len(queue) == 1 {
		cur := queue[0]
		queue = queue[1:]
		// 添加至拓扑序列
		res = append(res, cur)

		// 删除边
		for _, adj := range graph[cur] {
			inDegrees[adj]--
			if v, ok := inDegrees[adj]; ok && v == 0 {
				queue = append(queue, adj)
			}
		}
	}

	return reflect.DeepEqual(res, nums)
}
