package ofsort

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// Merge Sort

	// empty or single node
	if head == nil || head.Next == nil {
		return head
	}

	// split list
	left := head
	right := splitList(head)

	// recursion
	left = sortList(left)
	right = sortList(right)

	// merge
	return mergeList(left, right)
}

func splitList(head *ListNode) *ListNode {
	// fast start from second node
	// when fast reach tail
	// slow is the previous node of mid node
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// mid node
	right := slow.Next
	slow.Next = nil // break link

	return right
}

func mergeList(h1, h2 *ListNode) *ListNode {
	// sentinel node
	dummy := &ListNode{}
	cur := dummy

	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}

		cur = cur.Next
	}

	if h1 == nil {
		cur.Next = h2
	} else {
		cur.Next = h1
	}

	return dummy.Next
}
