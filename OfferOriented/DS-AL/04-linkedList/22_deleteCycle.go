package ordalist

func detectCycle(head *ListNode) *ListNode {
	// 获取环中节点
	inLoop := getNodeInLoop(head)
	if inLoop == nil {
		return nil
	}

	// 前后指针
	front := head
	// 相遇点即为环入口
	for front != inLoop {
		front = front.Next
		inLoop = inLoop.Next
	}

	return front
}

func getNodeInLoop(head *ListNode) *ListNode {
	// 边界: 空链表或单节点
	if head == nil || head.Next == nil {
		return nil
	}

	// 快慢指针
	slow := head.Next
	fast := slow.Next // 步数为 slow 的两倍
	for fast != nil {
		// 相遇点
		if slow == fast {
			return slow
		}
		// slow 走一步
		slow = slow.Next
		// fast 走两步
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return nil
}
