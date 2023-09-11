package ordalist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哨兵节点, 免去单独判断删除头节点的情况
	dummy := &ListNode{}
	dummy.Next = head

	// 前后双指针
	front, back := head, dummy
	// 前指针移动 n 步
	for i := 0; i < n; i++ {
		front = front.Next
	}

	// 移动双指针
	for front != nil {
		front = front.Next
		back = back.Next
	}

	// 删除倒数第n个
	back.Next = back.Next.Next
	return dummy.Next
}
