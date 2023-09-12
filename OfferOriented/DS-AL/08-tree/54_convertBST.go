package oftree

func convertBST(root *TreeNode) *TreeNode {
	st := make([]*TreeNode, 0) // stack
	cur := root
	sum := 0

	for cur != nil || len(st) > 0 {
		// right tree
		for cur != nil {
			st = append(st, cur)
			cur = cur.Right
		}

		// sum
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		sum += cur.Val
		cur.Val = sum

		// left tree
		cur = cur.Left
	}

	return root
}
