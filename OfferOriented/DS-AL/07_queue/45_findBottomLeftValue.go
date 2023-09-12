package ofqueue

func findBottomLeftValue(root *TreeNode) int {
	bottomVal := root.Val
	// 广度优先
	q1 := make([]*TreeNode, 0) // 当前层
	q2 := make([]*TreeNode, 0) // 下一层
	q1 = append(q1, root)
	for len(q1) > 0 {
		// 出队
		node := q1[0]
		q1 = q1[1:]
		// 下一层入队
		if node.Left != nil {
			q2 = append(q2, node.Left)
		}
		if node.Right != nil {
			q2 = append(q2, node.Right)
		}
		// 当前层结束
		if len(q1) == 0 {
			// 下一层
			q1, q2 = q2, q1
			if len(q1) > 0 {
				// 下一层最左
				bottomVal = q1[0].Val
			}
		}
	}

	return bottomVal
}
