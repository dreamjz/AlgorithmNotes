package oftrie

type MapSum struct {
	trie *MapTrie
}

func Constructor66() MapSum {
	return MapSum{trie: &MapTrie{}}
}

func (ms *MapSum) Insert(key string, val int) {
	ms.trie.Insert(key, val)
}

func (ms *MapSum) Sum(prefix string) int {
	cur := ms.trie
	// find last node in trie of prefix
	for _, ch := range prefix {
		c := ch - 'a'
		// not found
		if cur.children[c] == nil {
			return 0
		}

		cur = cur.children[c]
	}

	// calculate val in all subtree
	var dfsSum func(*MapTrie) int
	dfsSum = func(node *MapTrie) int {
		sum := 0
		if node == nil {
			return sum
		}

		// current node
		sum += node.val
		// subtree
		for _, child := range node.children {
			if child != nil {
				sum += dfsSum(child)
			}
		}

		return sum
	}

	return dfsSum(cur)
}

type MapTrie struct {
	children [26]*MapTrie
	val      int
}

func (mt *MapTrie) Insert(w string, val int) {
	cur := mt
	for _, ch := range w {
		c := ch - 'a'
		if cur.children[c] == nil {
			cur.children[c] = &MapTrie{}
		}

		cur = cur.children[c]
	}

	cur.val = val
}
