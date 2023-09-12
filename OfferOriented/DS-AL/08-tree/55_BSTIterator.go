package oftree

type BSTIterator struct {
	cur *TreeNode   // curren node
	st  []*TreeNode // stack
}

func Constructor55(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
		st:  make([]*TreeNode, 0),
	}
}

func (bi *BSTIterator) Next() int {
	// left tree
	for bi.cur != nil {
		bi.st = append(bi.st, bi.cur)
		bi.cur = bi.cur.Left
	}

	bi.cur = bi.st[len(bi.st)-1]
	bi.st = bi.st[:len(bi.st)-1]
	val := bi.cur.Val

	// right tree
	bi.cur = bi.cur.Right

	return val
}

func (bi *BSTIterator) HasNext() bool {
	return bi.cur != nil || len(bi.st) > 0
}
