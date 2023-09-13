package oftrie

type Trie struct {
	children [26]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, c := range word {
		c -= 'a'
		if cur.children[c] == nil {
			cur.children[c] = &Trie{}
		}
		cur = cur.children[c]
	}
	cur.isWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, c := range word {
		c -= 'a'
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}

	return cur.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, c := range prefix {
		c -= 'a'
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}

	return true
}
