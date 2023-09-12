package oftree

func pathSum(root *TreeNode, targetSum int) int {
	m := make(map[int]int)
	// 初始值
	m[0] = 1
	return dfs50(root, m, targetSum, 0)
}

func dfs50(root *TreeNode, m map[int]int, sum, path int) int {
	// 递归出口
	if root == nil {
		return 0
	}
	// 当前路径和
	path += root.Val
	// 寻找目标路径个数
	cnt := 0
	if v, ok := m[path-sum]; ok {
		cnt = v
	}

	// 记录次数
	m[path]++

	// 左右子树
	cnt += dfs50(root.Left, m, sum, path)
	cnt += dfs50(root.Right, m, sum, path)

	// 去除已经统计的路径
	m[path]--

	return cnt
}
