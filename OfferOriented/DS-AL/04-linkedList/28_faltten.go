package ordalist

type Node struct {
	Val   int
	Next  *Node
	Prev  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	flattenGetTail(root)
	return root
}

func flattenGetTail(head *Node) *Node {
	node := head
	var tail *Node // 展开后的尾节点

	// 遍历链表
	for node != nil {
		next := node.Next
		// 展开子链表
		if node.Child != nil {
			child := node.Child                // 子链表头节点
			childTail := flattenGetTail(child) // 递归展开子链表, 返回展开后的尾节点

			// 插入原链表
			node.Child = nil // 断开子链表
			node.Next = child
			child.Prev = node
			childTail.Next = next
			if next != nil { // 注意双向链表的指针有两个
				next.Prev = childTail
			}
			// 更新当前展开的尾节点
			tail = childTail
		} else { // 递归出口, 无子链表
			tail = node
		}
		// 更新当前节点
		node = next
	}

	return tail
}
