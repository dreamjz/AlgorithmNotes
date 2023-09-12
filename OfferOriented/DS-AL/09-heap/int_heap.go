package ofheap

// 最小堆
type minIntHeap []int

func (mh minIntHeap) Len() int {
	return len(mh)
}

func (mh minIntHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minIntHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minIntHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *minIntHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

// 最大堆
type maxIntHeap []int

func (mh maxIntHeap) Len() int {
	return len(mh)
}

func (mh maxIntHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh maxIntHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxIntHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *maxIntHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}
