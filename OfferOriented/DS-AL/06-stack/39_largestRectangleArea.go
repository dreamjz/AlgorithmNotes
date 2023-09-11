package ofstack

func largestRectangleAreaBruteForce(heights []int) int {
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			minHeight = min(minHeight, heights[j])
			area := (j - i + 1) * minHeight
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func largestRectangleAreaRecursion(heights []int) int {
	return largestArea(heights, 0, len(heights))
}

func largestArea(heights []int, start, end int) int {
	// 递归出口
	if start == end {
		return 0
	}
	if start+1 == end {
		return heights[start]
	}

	// 寻找最低高度
	minIdx := start
	for i := start + 1; i < end; i++ {
		if heights[i] < heights[minIdx] {
			minIdx = i
		}
	}
	minHeight := heights[minIdx]
	area := (end - start) * minHeight            // 中间
	left := largestArea(heights, start, minIdx)  // 左边
	right := largestArea(heights, minIdx+1, end) // 右边

	area = max(area, left)
	return max(area, right)
}

func largestRectangleAreaStack(heights []int) int {
	stack := make([]int, 0, len(heights))
	// 假设最左边有个的柱子, 方便计算最矮柱子对应的最大矩形
	stack = append(stack, -1)

	maxArea := 0
	for i := range heights {
		// 出栈条件
		for stack[len(stack)-1] != -1 && heights[i] <= heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[idx]
			width := i - stack[len(stack)-1] - 1
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		h := heights[i]
		w := len(heights) - stack[len(stack)-1] - 1
		maxArea = max(maxArea, h*w)
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
