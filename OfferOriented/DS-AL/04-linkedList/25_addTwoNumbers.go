package ordalist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 反转链表
	rl1 := reverseList25(l1)
	rl2 := reverseList25(l2)
	// 按位计算
	res := addList(rl1, rl2)

	return reverseList25(res)
}

func reverseList25(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func addList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	carry := 0 // 进位
	node := dummy
	for l1 != nil || l2 != nil {
		sum := 0
		sum += getVal(l1) + getVal(l2) + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		// 记录结果
		v := sum % 10
		node.Next = &ListNode{Val: v}
		node = node.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		node.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}

func getVal(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}
