package ordalist

func reorderList(head *ListNode) {
	dummy := &ListNode{Next: head}
	// 快慢指针, 寻找中间节点
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 分离前后半链表
	tmp := slow.Next
	slow.Next = nil

	link(head, reverseList26(tmp))
}

func reverseList26(head *ListNode) *ListNode {
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

func link(node1, node2 *ListNode) {
	dummy := &ListNode{}

	node := dummy
	for node1 != nil && node2 != nil {
		tmp := node1.Next

		node.Next = node1
		node1.Next = node2
		node = node2

		node1 = tmp
		node2 = node2.Next
	}
	if node1 != nil {
		node.Next = node1
	}
}
