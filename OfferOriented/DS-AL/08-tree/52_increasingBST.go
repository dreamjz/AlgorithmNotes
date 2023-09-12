package oftree

func increasingBST(root *TreeNode) *TreeNode {
	st := make([]*TreeNode, 0) // stack
	cur := root                // current node
	var prev *TreeNode         // previous node
	var newRoot *TreeNode      // new tree root

	// inorder dsf
	for cur != nil || len(st) > 0 {
		// left
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		// mid
		cur = st[len(st)-1]
		st = st[:len(st)-1]

		// 连接到前一节点
		if prev != nil {
			prev.Right = cur
		} else {
			// 新树的根节点
			newRoot = cur
		}

		// right
		prev = cur
		// 断开左子树
		cur.Left = nil
		cur = cur.Right
	}

	return newRoot
}
