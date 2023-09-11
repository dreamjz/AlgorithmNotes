package ofstack

func dailyTemperatures(temperatures []int) []int {
	// 栈
	stack := make([]int, 0, len(temperatures))

	result := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] { // 出栈条件
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev // 天数差
		}

		stack = append(stack, i)
	}

	return result
}
