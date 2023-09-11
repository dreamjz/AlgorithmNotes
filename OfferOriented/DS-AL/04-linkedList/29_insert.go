package ordalist

type Node29 struct {
	Val  int
	Next *Node29
}

func insert(aNode *Node29, x int) *Node29 {
	newNode := &Node29{Val: x}
	// 边界条件
	if aNode == nil { // 空链表
		newNode.Next = newNode
		return newNode
	} else if aNode.Next == nil { // 只有一个节点
		aNode.Next = newNode
		newNode.Next = aNode
	} else {
		insertCore(aNode, newNode)
	}

	return aNode
}

func insertCore(head, node *Node29) {
	cur := head
	next := cur.Next
	biggest := cur

	// 寻找插入位置
	for !(cur.Val <= node.Val && next.Val >= node.Val) && next != head {
		cur = next
		next = cur.Next
		if cur.Val >= biggest.Val {
			biggest = cur
		}
	}

	// 插入
	if cur.Val <= node.Val && next.Val >= node.Val { // 插入位置在中间
		cur.Next = node
		node.Next = next
	} else { // 插入位置为末尾
		node.Next = biggest.Next
		biggest.Next = node
	}

}
