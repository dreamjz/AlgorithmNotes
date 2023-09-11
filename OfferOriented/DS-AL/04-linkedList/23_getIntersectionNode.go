package ordalist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 获取链表长度
	lenA := countList(headA)
	lenB := countList(headB)

	// 前后双指针
	longer, shorter := &ListNode{}, &ListNode{}
	delta := 0
	if lenA > lenB {
		longer = headA
		shorter = headB
		delta = lenA - lenB
	} else {
		longer = headB
		shorter = headA
		delta = lenB - lenA
	}

	// 先移动前指针
	for i := 0; i < delta; i++ {
		longer = longer.Next
	}
	// 同步移动
	for longer != shorter {
		longer = longer.Next
		shorter = shorter.Next
	}

	return longer
}

func countList(head *ListNode) int {
	count := 0
	for head != nil {
		head = head.Next
		count++
	}

	return count
}
