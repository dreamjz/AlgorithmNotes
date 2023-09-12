package oftree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var f func(*TreeNode) string
	f = func(node *TreeNode) string {
		var sb strings.Builder
		// 递归出口
		if node == nil {
			return "null"
		}
		// root
		rootStr := strconv.Itoa(node.Val)
		// left
		leftStr := f(node.Left)
		// right
		rightStr := f(node.Right)

		sb.WriteString(rootStr)
		sb.WriteByte(',')
		sb.WriteString(leftStr)
		sb.WriteByte(',')
		sb.WriteString(rightStr)

		return sb.String()
	}
	return f(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	return dfs(&strs)
}

func dfs(strs *[]string) *TreeNode {
	if len(*strs) == 0 {
		return nil
	}
	// node val
	str := (*strs)[0]
	*strs = (*strs)[1:]
	if str == "null" {
		return nil
	}

	val, _ := strconv.Atoi(str)
	// root node
	root := &TreeNode{Val: val}
	// left tree
	root.Left = dfs(strs)
	// right tree
	root.Right = dfs(strs)

	return root
}
