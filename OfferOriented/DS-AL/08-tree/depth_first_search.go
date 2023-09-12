package oftree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	dfsInorderRecursively(root, &res)
	return res
}

func dfsInorderRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		dfsInorderRecursively(root.Left, res)
		*res = append(*res, root.Val)
		dfsInorderRecursively(root.Right, res)
	}
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	// 栈
	st := make([]*TreeNode, 0)
	// 当前节点
	cur := root
	for cur != nil || len(st) != 0 {
		// 左子树
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		// 根
		cur = st[len(st)-1]
		st = st[0 : len(st)-1]
		res = append(res, cur.Val)
		// 右子树
		cur = cur.Right
	}

	return res
}

func preorderTraversalWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	dfsPreorderRecursively(root, &res)
	return res
}

func dfsPreorderRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		dfsPreorderRecursively(root.Left, res)
		dfsPreorderRecursively(root.Right, res)
	}
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 栈
	st := make([]*TreeNode, 0)
	// 当前遍历节点
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil {
			// 根
			res = append(res, cur.Val)
			// 左子树
			st = append(st, cur)
			cur = cur.Left
		}
		// 右子树
		cur = st[len(st)-1]
		st = st[0 : len(st)-1]
		cur = cur.Right
	}

	return res
}

func postorderTraversalWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	dfsInorderRecursively(root, &res)
	return res
}

func dfsPostorderRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		dfsInorderRecursively(root.Left, res)
		dfsInorderRecursively(root.Right, res)
		*res = append(*res, root.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// stack
	st := make([]*TreeNode, 0)
	// current node
	cur := root
	// previous node
	var prev *TreeNode
	for cur != nil || len(st) > 0 {
		// left tree
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		if cur.Right != nil && cur.Right != prev {
			// right tree
			cur = cur.Right
		} else {
			// root
			res = append(res, cur.Val)
			st = st[:len(st)-1]
			prev = cur
			cur = nil
		}
	}

	return res
}
