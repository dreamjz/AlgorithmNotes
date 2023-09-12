package oftree

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	// 递归函数
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 计算左右子树最大路径和
		leftMax := max(dfs(node.Left), 0)
		rightMax := max(dfs(node.Right), 0)

		// 当前路径和
		path := node.Val + leftMax + rightMax

		// 更新最大
		maxSum = max(maxSum, path)

		return node.Val + max(leftMax, rightMax)
	}
	dfs(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
