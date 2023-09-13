package oftrie

func findMaximumXOR(nums []int) int {
	// 构建前缀树
	trie := buildBitTrie(nums)

	// 计算最大异或值
	maxXOR := 0
	for _, n := range nums {
		cur := trie
		xor := 0
		for i := 31; i >= 0; i-- {
			bit := (n >> i) & 1
			if cur.children[1-bit] != nil { // 相反bit
				xor = (xor << 1) + 1 // 异或结果 1
				cur = cur.children[1-bit]
			} else { // 相同 bit
				xor <<= 1
				cur = cur.children[bit]
			}
		}

		maxXOR = max(maxXOR, xor)
	}

	return maxXOR
}

func buildBitTrie(nums []int) *BitTrie {
	root := &BitTrie{}
	for _, n := range nums {
		root.Insert(int32(n))
	}
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type BitTrie struct {
	children [2]*BitTrie
}

func (bt *BitTrie) Insert(n int32) {
	cur := bt
	for i := 31; i >= 0; i-- {
		bit := (n >> i) & 1
		if cur.children[bit] == nil {
			cur.children[bit] = &BitTrie{}
		}

		cur = cur.children[bit]
	}
}
