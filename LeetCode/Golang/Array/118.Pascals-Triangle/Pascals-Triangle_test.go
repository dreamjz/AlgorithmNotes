package array

import (
	"fmt"
	"testing"
)

const (
	errFormat  = "Wrong Answer: [input]: %v, [output]: %v, [answer]: %v\n"
	passFormat = "Pass: [input]: %v, [output]: %v, [answer]: %v\n"
)

type Question118 struct {
	param int
	ans   [][]int
}

var (
	qs []Question118
)

func init() {
	qs = []Question118{
		{
			param: 5,
			ans:   [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			param: 1,
			ans:   [][]int{{1}},
		},
	}
}

func TestGenerate(t *testing.T) {
	for _, q := range qs {
		res := generate(q.param)
		if !intTowDimensionSliceEq(q.ans, res) {
			t.Errorf(errFormat, q.param, res, q.ans)
			return
		}
		fmt.Printf(passFormat, q.param, res, q.ans)

	}
}

func intSliceEq(a, b []int) bool {
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

func intTowDimensionSliceEq(a, b [][]int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !intSliceEq(a[i], b[i]) {
			return false
		}
	}
	return true
}
