package ofqueue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	queue []*TreeNode
	root  *TreeNode
}

func Constructor43(root *TreeNode) CBTInserter {
	// 使用广度优先搜索, 查找不含左/右子节点的节点
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// 广度优先搜索
	for len(queue) > 0 && queue[0].Left != nil && queue[0].Right != nil {
		// 出列
		node := queue[0]
		queue = queue[1:]
		// 添加左右节点
		queue = append(queue, node.Left, node.Right)
	}

	return CBTInserter{
		queue: queue,
		root:  root,
	}
}

func (ci *CBTInserter) Insert(v int) int {
	node := &TreeNode{Val: v}
	// 获取最左边的有空缺的节点
	parent := ci.queue[0]
	if parent.Left == nil { // 插入左边
		parent.Left = node
	} else { // 插入右边
		parent.Right = node
		// 此时左右节点都存在, 出队
		ci.queue = ci.queue[1:]
		// 子节点入队
		ci.queue = append(ci.queue, parent.Left, parent.Right)
	}

	return parent.Val
}

func (ci *CBTInserter) GetRoot() *TreeNode {
	return ci.root
}
