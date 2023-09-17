package ofgraph

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		set[n] = true
	}

	var bfs func(map[int]bool, int) int
	bfs = func(set map[int]bool, n int) int {
		q := []int{n}
		delete(set, n)

		l := 1
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			adjs := []int{cur - 1, cur + 1}
			for _, adj := range adjs {
				if set[adj] {
					q = append(q, adj)
					delete(set, adj)
					l++
				}
			}
		}

		return l
	}

	res := 0
	for _, n := range nums {
		if len(set) > 0 {
			res = max(res, bfs(set, n))
		}
	}

	return res
}
func longestConsecutive2(nums []int) int {
	// 并查集
	n := len(nums)
	fathers := make(map[int]int, n)
	cnts := make(map[int]int, n) // 子集长度
	set := make(map[int]bool, n) // 数字集合
	for _, n := range nums {
		fathers[n] = n
		cnts[n] = 1
		set[n] = true
	}

	// 合并
	for _, n := range nums {
		if set[n-1] {
			union119(fathers, cnts, n, n-1)
		}

		if set[n+1] {
			union119(fathers, cnts, n, n+1)
		}
	}

	res := 0
	for _, v := range cnts {
		if res < v {
			res = v
		}
	}

	return res
}

func findFathers(fathers map[int]int, i int) int {
	if fathers[i] != i {
		fathers[i] = findFathers(fathers, fathers[i])
	}
	return fathers[i]
}

func union119(fathers, cnts map[int]int, i, j int) {
	fI := findFathers(fathers, i)
	fJ := findFathers(fathers, j)
	if fI != fJ {
		fathers[fI] = fJ
		cnts[fJ] = cnts[fI] + cnts[fJ]
	}
}
