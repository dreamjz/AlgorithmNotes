package Array

import (
	"fmt"
	"testing"
)

type question344 struct {
	param string
	ans   string
}

func TestReverseString(t *testing.T) {
	qs := []question344{
		{
			param: "leetcode",
			ans:   "edocteel",
		},
		{
			param: "hello",
			ans:   "olleh",
		},
		{
			param: "",
			ans:   "",
		},
	}

	fmt.Println("------------------------ Solution 1------------------------")
	for _, q := range qs {
		in := []byte(q.param)
		reverseString(in)
		res := string(in)
		if q.ans != res {
			t.Errorf("Wrong anwser: [Input]: %v, [Output]: %v, [Anwser]: %v\n, ", q.param, res, q.ans)
		}
		fmt.Printf("[Input]: %v, [Output]: %v, [Anwser]: %v\n", q.param, res, q.ans)
	}
}
