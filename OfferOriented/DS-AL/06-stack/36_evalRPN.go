package ofstack

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		num, err := strconv.Atoi(t)
		if err == nil {
			// 数字入栈
			stack = append(stack, num)
		} else { // 符号, 则数字出栈
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch t {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		}
	}

	return stack[len(stack)-1]
}
