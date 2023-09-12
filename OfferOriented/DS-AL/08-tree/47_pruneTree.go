package oftree

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 剪除左子树
	root.Left = pruneTree(root.Left)
	// 剪除右子树
	root.Right = pruneTree(root.Right)
	// 剪除当前树
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
