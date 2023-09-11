package ordalist

func isPalindrome(head *ListNode) bool {
	// 边界条件
	if head == nil || head.Next == nil {
		return true
	}

	dummy := &ListNode{Next: head}
	// 快慢双指针, 寻找中间节点
	slow, fast := dummy, dummy.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 前半部分
	secondHalf := &ListNode{}
	if fast == nil {
		secondHalf = slow.Next
	} else if fast.Next == nil { // 链表长为奇数
		secondHalf = slow.Next.Next
	}

	// 前后分离
	slow.Next = nil
	// 反转后半
	secondHalf = reverseList27(secondHalf)
	// 比较
	return listEqual(head, secondHalf)
}

func reverseList27(head *ListNode) *ListNode {
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

func listEqual(head1, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return head1 == nil && head2 == nil
}
