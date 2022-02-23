package array

import (
	"fmt"
	"testing"
)

type Question151 struct {
	params string
	ans    string
}

func TestReverseWords1(t *testing.T) {
	qs := []Question151{
		{
			params: "the sky is blue",
			ans:    "blue is sky the",
		},
		{
			params: "the sky   is blue",
			ans:    "blue is sky the",
		},
		{
			params: "  the sky is blue  ",
			ans:    "blue is sky the",
		},
	}

	fmt.Println("------------------------Solution 1------------------------")
	for _, q := range qs {
		res := reverseWords1(q.params)
		if q.ans != res {
			t.Errorf("Wrong anwser: [Input]: %v, [Output]: %v, [Anwser]: %v\n, ", q.params, res, q.ans)
		}
		fmt.Printf("[Input]: %v, [Output]: %v, [Anwser]: %v\n", q.params, res, q.ans)
	}
	fmt.Println("------------------------Solution 2------------------------")
	for _, q := range qs {
		res := reverseWords2(q.params)
		if q.ans != res {
			t.Errorf("Wrong anwser: [Input]: %v, [Output]: %v, [Anwser]: %v\n, ", q.params, res, q.ans)
		}
		fmt.Printf("[Input]: %v, [Output]: %v, [Anwser]: %v\n", q.params, reverseWords2(q.params), q.ans)
	}
	fmt.Println("------------------------Solution 3------------------------")
	for _, q := range qs {
		res := reverseWords3(q.params)
		if q.ans != res {
			t.Errorf("Wrong anwser: [Input]: %v, [Output]: %v, [Anwser]: %v\n, ", q.params, res, q.ans)
		}
		fmt.Printf("[Input]: %v, [Output]: %v, [Anwser]: %v\n", q.params, reverseWords2(q.params), q.ans)
	}
}
