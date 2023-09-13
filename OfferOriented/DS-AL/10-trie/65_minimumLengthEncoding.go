package oftrie

func minimumLengthEncoding(words []string) int {
	// 构建前缀树
	reTrie := buildReTrie(words)

	// 深度优先遍历
	// 路径长度和
	sum := 0
	var dfs func(*ReTrie, int)
	dfs = func(node *ReTrie, length int) {
		isLeaf := true
		for _, child := range node.children {
			if child != nil {
				isLeaf = false
				// 遍历子节点
				dfs(child, length+1)
			}
		}

		// 到达叶节点, 统计路径长度
		if isLeaf {
			sum += length
		}
	}
	// 单词后多一个'#'
	dfs(reTrie, 1)

	return sum
}

func buildReTrie(words []string) *ReTrie {
	root := &ReTrie{}
	for _, w := range words {
		root.Insert(w)
	}
	return root
}

type ReTrie struct {
	children [26]*ReTrie
}

func (rt *ReTrie) Insert(w string) {
	cur := rt
	for i := len(w) - 1; i >= 0; i-- {
		c := w[i] - 'a'
		if cur.children[c] == nil {
			cur.children[c] = &ReTrie{}
		}

		cur = cur.children[c]
	}
}
