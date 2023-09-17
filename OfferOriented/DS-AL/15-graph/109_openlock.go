package ofgraph

import "strings"

func openLock(deadends []string, target string) int {
	init := "0000" // 初始

	// 锁定 哈希表
	mDead := make(map[string]bool)
	for _, d := range deadends {
		mDead[d] = true
	}

	// 边界
	if mDead[init] || mDead[target] {
		return -1
	}

	// 已访问状态 哈希表
	mVisited := make(map[string]bool)

	// 双队列
	queue1 := []string{init}
	queue2 := []string{}

	steps := 0 // 路径长
	// BFS
	for len(queue1) > 0 {
		// 出队
		cur := queue1[0]
		queue1 = queue1[1:]
		// 解锁
		if cur == target {
			return steps
		}

		// 相邻节点入队
		adjs := getAdjs(cur)
		for _, adj := range adjs {
			// 不是锁定码 且 未访问
			if !mDead[adj] && !mVisited[adj] {
				queue2 = append(queue2, adj)
				mVisited[adj] = true
			}
		}

		// 下一层
		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = []string{}
			steps++
		}
	}

	return -1
}

func getAdjs(s string) []string {
	// 共有 8 种新状态
	adjs := make([]string, 0, 8)

	for i, c := range s {
		var sb strings.Builder
		sb.Grow(4)

		var ch rune
		// 增加
		if c == '9' {
			ch = '0'
		} else {
			ch = c + 1
		}
		addAdjs(&sb, &adjs, s, ch, i)

		// 减少
		if c == '0' {
			ch = '9'
		} else {
			ch = c - 1
		}
		addAdjs(&sb, &adjs, s, ch, i)
	}

	return adjs
}

func addAdjs(sb *strings.Builder, adjs *[]string, s string, ch rune, i int) {
	sb.WriteString(s[:i])
	sb.WriteRune(ch)
	sb.WriteString(s[i+1:])
	*adjs = append(*adjs, sb.String())
	// Reset builder
	sb.Reset()
}
