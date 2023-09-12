package ofheap

import "container/heap"

type KthLargest struct {
	minHeap *minIntHeap
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	mh := &minIntHeap{}
	kl := &KthLargest{
		minHeap: mh,
		size:    k,
	}

	for _, n := range nums {
		kl.Add(n)
	}

	return *kl
}

func (kl *KthLargest) Add(val int) int {
	if kl.minHeap.Len() < kl.size {
		heap.Push(kl.minHeap, val)
	} else if val > (*kl.minHeap)[0] {
		heap.Pop(kl.minHeap)
		heap.Push(kl.minHeap, val)
	}

	return (*kl.minHeap)[0]
}
