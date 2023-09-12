package oftree

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	cur := root // current node
	for cur != nil {
		if cur.Val > p.Val {
			// left tree
			// record result
			res = cur
			cur = cur.Left
		} else {
			// right tree
			cur = cur.Right
		}
	}

	return res
}
