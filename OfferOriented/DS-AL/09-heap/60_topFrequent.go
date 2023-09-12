package ofheap

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	res := []int{}

	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	pq := &PriorityQueue{}
	for n, fre := range m {
		nf := &numFre{
			val: n,
			fre: fre,
		}

		if pq.Len() < k {
			heap.Push(pq, nf)
		} else if nf.fre > (*pq)[0].fre {
			heap.Pop(pq)
			heap.Push(pq, nf)
		}
	}

	for _, nf := range *pq {
		res = append(res, nf.val)
	}
	return res
}

type numFre struct {
	val   int
	fre   int
	index int
}

type PriorityQueue []*numFre

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].fre < pq[j].fre
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	idx := len(*pq)
	nf := x.(*numFre)
	nf.index = idx
	*pq = append(*pq, nf)
}

func (pq *PriorityQueue) Pop() any {
	x := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil // avoid memory leak
	x.index = -1            // for safety
	*pq = (*pq)[:len(*pq)-1]
	return x
}
