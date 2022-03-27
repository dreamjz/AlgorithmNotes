package array

import (
	"fmt"
	"log"
	"testing"
)

type Question1 struct {
	params Param
	ans    []int
}

type Param struct {
	nums   []int
	target int
}

var (
	qs1 []Question1
)

const (
	errFormat  = "Wrong Answer: [input]: %v, [output]: %v, [answer]: %v\n"
	passFormat = "Pass: [input]: %v, [output]: %v, [answer]: %v\n"
)

func init() {
	qs1 = []Question1{
		{
			params: Param{[]int{2, 7, 11, 15}, 9},
			ans:    []int{0, 1},
		},
		{
			params: Param{[]int{3, 2, 4}, 6},
			ans:    []int{1, 2},
		},
		{
			params: Param{[]int{3, 3}, 6},
			ans:    []int{0, 1},
		},
	}
}

func TestBruteForce(t *testing.T) {
	log.Println("------------------------ Brute Force ------------------------")
	for _, q := range qs1 {
		res := bruteForce(q.params.nums, q.params.target)
		if !intsEqWithoutOrder(res, q.ans) {
			t.Errorf(errFormat, q.params, res, q.ans)
		}
		fmt.Printf(passFormat, q.params, res, q.ans)
	}
}

func TestHashTable(t *testing.T) {
	log.Println("------------------------ Hash Table ------------------------")
	for _, q := range qs1 {
		res := hashTable(q.params.nums, q.params.target)
		if !intsEqWithoutOrder(res, q.ans) {
			t.Errorf(errFormat, q.params, res, q.ans)
		}
		fmt.Printf(passFormat, q.params, res, q.ans)
	}
}

func intsEqWithoutOrder(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	hashMap := map[int]int{}
	for i, v := range a {
		hashMap[v] = i
	}
	for _, v := range b {
		if _, ok := hashMap[v]; !ok {
			return false
		}
	}
	return true
}
