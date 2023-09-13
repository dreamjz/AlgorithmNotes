package oftrie

type MagicDictionary struct {
	trie *Trie
}

func Constructor64() MagicDictionary {
	return MagicDictionary{trie: &Trie{}}
}

func (md *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		md.trie.Insert(w)
	}
}

func (md *MagicDictionary) Search(searchWord string) bool {
	return dfs(md.trie, searchWord, false)
}

func dfs(node *Trie, s string, modified bool) bool {
	// 查找结束
	if s == "" {
		return node.isWord && modified
	}

	// 深度优先
	// 当前节点相同则查找子节点
	c := s[0] - 'a'
	if node.children[c] != nil && dfs(node.children[c], s[1:], modified) {
		return true
	}

	// 未修改则用任意子节点替换, 继续查找
	if !modified {
		for i, child := range node.children {
			if i != int(c) && child != nil && dfs(child, s[1:], true) {
				return true
			}
		}
	}

	return false
}
