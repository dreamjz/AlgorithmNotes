package oftree

func sumNumbers(root *TreeNode) int {
	// 递归
	return dfs49(root, 0)
}

func dfs49(root *TreeNode, path int) int {
	// 递归出口
	if root == nil {
		return 0
	}

	path = path*10 + root.Val
	// 叶节点
	if root.Left == nil && root.Right == nil {
		return path
	}

	return dfs49(root.Left, path) + dfs49(root.Right, path)
}
