package oftree

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool, 0)

	st := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(st) > 0 {
		// left
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}

		// node
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if v, ok := m[k-cur.Val]; ok {
			return v
		}
		m[cur.Val] = true

		// right
		cur = cur.Right
	}

	return false
}
