package Array

import (
	"fmt"
	"testing"
)

type question167 struct {
	param param167
	ans   []int
}

type param167 struct {
	param1 []int
	param2 int
}

func TestTwoSum1(t *testing.T) {
	qs := []question167{
		{
			param: param167{
				param1: []int{2, 7, 11, 15},
				param2: 9,
			},
			ans: []int{1, 2},
		},
	}

	fmt.Println("------------------------ Solution 1 ------------------------")
	for _, q := range qs {
		res := twoSum1(q.param.param1, q.param.param2)
		if !intsEq(res, q.ans) {
			t.Errorf("Wrong anwser: [Input]: %v, [Output]: %v, [Anwser]: %v\n, ", q.param, res, q.ans)
		}
		fmt.Printf("[Input]: %v, [Output]: %v, [Anwser]: %v\n", q.param, res, q.ans)
	}
}

func intsEq(a, b []int) bool {
	// golang: bool != boll 可以表示异或
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
