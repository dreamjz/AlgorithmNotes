package ofsort

func mergeKLists(lists []*ListNode) *ListNode {
	// empty
	if len(lists) == 0 {
		return nil
	}
	// single node
	if len(lists) == 1 {
		return lists[0]
	}

	// split lists
	mid := len(lists) >> 1
	left := lists[:mid]
	right := lists[mid:]

	// merge
	return mergeList(mergeKLists(left), mergeKLists(right))
}
