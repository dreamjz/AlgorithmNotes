package ofstack

func asteroidCollision(asteroids []int) []int {
	// 栈
	stack := make([]int, 0)

	for _, a := range asteroids {
		alive := true
		for alive && a < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = stack[len(stack)-1] < -a // 当前元素是否存活
			if stack[len(stack)-1] <= -a {   // 栈顶爆炸
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, a)
		}
	}

	return stack
}
