package ofqueue

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	// 边界
	if root == nil {
		return result
	}

	// 广度优先
	q1 := make([]*TreeNode, 0) // 当前层
	q2 := make([]*TreeNode, 0) // 下一层
	q1 = append(q1, root)
	for len(q1) > 0 {
		// 出队
		node := q1[0]
		q1 = q1[1:]
		// 下层入队
		if node.Left != nil {
			q2 = append(q2, node.Left)
		}
		if node.Right != nil {
			q2 = append(q2, node.Right)
		}
		// 当前层结束
		if len(q1) == 0 {
			// 最右侧节点
			result = append(result, node.Val)
			// 开始下一层
			q1, q2 = q2, q1
		}
	}

	return result
}
