package oftrie

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	// 构建前缀树
	trie := buildTrie(dictionary)

	// 替换单词
	var sb strings.Builder
	sb.Grow(len(sentence))
	words := strings.Split(sentence, " ")
	for i, w := range words {
		// 寻找前缀
		prefix := findPrefix(trie, w)
		if prefix != "" {
			w = prefix
		}

		sb.WriteString(w)
		if i != len(words)-1 {
			sb.WriteByte(' ')
		}
	}

	return sb.String()
}

func findPrefix(root *Trie, w string) string {
	prefix := ""

	var sb strings.Builder
	sb.Grow(len(w))
	cur := root
	for _, c := range w {
		ch := c
		c -= 'a'
		// 已找到或者不存在
		if cur.isWord || cur.children[c] == nil {
			break
		}
		sb.WriteRune(ch)

		cur = cur.children[c]
	}

	// 已找到
	if cur.isWord {
		prefix = sb.String()
	}

	return prefix
}

func buildTrie(dic []string) *Trie {
	root := &Trie{}
	for _, w := range dic {
		root.Insert(w)
	}

	return root
}
