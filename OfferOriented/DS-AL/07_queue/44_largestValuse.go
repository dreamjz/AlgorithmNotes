package ofqueue

import "math"

func largestValuesWithSingleQ(root *TreeNode) []int {
	result := make([]int, 0)
	// 边界条件
	if root == nil {
		return result
	}

	// 当前和下一层的节点数
	cur, next := 0, 0
	// 每层最大节点
	maxVal := math.MinInt
	// 广度优先搜索
	q := make([]*TreeNode, 0)
	q = append(q, root)
	cur++
	for len(q) > 0 {
		// 当前层节点出队
		node := q[0]
		q = q[1:]
		maxVal = max(maxVal, node.Val)
		cur--
		// 添加下一层
		if node.Left != nil {
			q = append(q, node.Left)
			next++
		}
		if node.Right != nil {
			q = append(q, node.Right)
			next++
		}
		// 当前层结束
		if cur == 0 {
			result = append(result, maxVal)
			maxVal = math.MinInt
			cur = next
			next = 0
		}
	}

	return result
}

func largestValuesWithDoubleQ(root *TreeNode) []int {
	result := make([]int, 0)
	// 边界
	if root == nil {
		return result
	}
	// 双队列
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)
	// 最大值
	maxVal := math.MinInt
	// 广度优先遍历
	q1 = append(q1, root)
	for len(q1) > 0 {
		// 出队
		node := q1[0]
		q1 = q1[1:]
		maxVal = max(maxVal, node.Val)
		// 下一层入队
		if node.Left != nil {
			q2 = append(q2, node.Left)
		}
		if node.Right != nil {
			q2 = append(q2, node.Right)
		}
		// 当前层结束
		if len(q1) == 0 {
			result = append(result, maxVal)
			maxVal = math.MinInt
			// 开始下一层
			q1, q2 = q2, q1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
