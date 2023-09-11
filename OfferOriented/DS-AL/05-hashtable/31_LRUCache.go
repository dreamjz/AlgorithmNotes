package orhash

type LRUCache struct {
	cache    map[int]*ListNode31
	Capacity int
	/* 哨兵节点 */
	head *ListNode31
	tail *ListNode31
}

type ListNode31 struct {
	Key  int
	Val  int
	Prev *ListNode31
	Next *ListNode31
}

func (lru *LRUCache) MoveToTail(node *ListNode31, val int) {
	lru.DeleteNode(node)
	node.Val = val
	lru.InsertToTail(node)
}

func (lru *LRUCache) DeleteNode(node *ListNode31) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) InsertToTail(node *ListNode31) {
	lru.tail.Prev.Next = node
	node.Prev = lru.tail.Prev
	node.Next = lru.tail
	lru.tail.Prev = node
}

func ConstructorLRU(capacity int) LRUCache {
	head := &ListNode31{Key: -1, Val: -1}
	tail := &ListNode31{Key: -1, Val: -1}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cache:    make(map[int]*ListNode31, capacity),
		Capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	lru.MoveToTail(node, node.Val) // 移动至表尾
	return node.Val
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.MoveToTail(node, value)
	} else {
		if len(lru.cache) == lru.Capacity {
			/* 移除头节点 */
			toBeDeleted := lru.head.Next
			lru.DeleteNode(toBeDeleted)
			delete(lru.cache, toBeDeleted.Key)
		}

		newNode := &ListNode31{Key: key, Val: value}
		lru.InsertToTail(newNode)
		lru.cache[key] = newNode
	}
}
