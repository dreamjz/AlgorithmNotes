package ordalist

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		// 反转
		cur.Next = prev
		// 更新
		prev = cur
		cur = next
	}

	return prev
}
