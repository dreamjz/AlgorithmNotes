package ordastring

import "math"

func minWindow(s string, t string) string {
	minStr := ""

	// 边界
	if len(s) < len(t) {
		return minStr
	}

	// 构建哈希表, 字符:次数, 长度为英文字符个数
	m := make(map[byte]int, 52)

	// 扫描 t
	for i := range t {
		m[t[i]]++
	}

	// 扫描 s, 移动子区间
	left, right := 0, 0       // 当前子区间
	minLen := math.MaxInt     // 最小子串长度
	minLeft, minRight := 0, 0 // 最小子串区间
	count := len(m)           // 区间中未包含的t中字符的个数

	// 边界值: 右指针已经到达末尾, 此时刚好包含所有的t中字符
	// 需要进行最后一次计算
	for right < len(s) || (count == 0 && right == len(s)) {
		// 存在未包含的t中字符, 需要扩大区间
		if count > 0 {
			// 包含t中字符
			if _, ok := m[s[right]]; ok {
				// 减少次数
				m[s[right]]--
				// 已包含全部的当前字符
				if m[s[right]] == 0 {
					count--
				}
			}
			// 增大区间
			right++
		} else { // 已包含所有的t中字符, 减小区间
			// 记录最小区间
			// s[l, r] 长度为 r-l
			if (right - left) < minLen {
				minLen = right - left
				minLeft = left
				minRight = right
			}

			if _, ok := m[s[left]]; ok {
				// 增加次数
				m[s[left]]++
				// 存在不包含的t中字符
				if m[s[left]] == 1 {
					count++
				}
			}
			// 减小区间
			left++
		}
	}

	// 存在
	if minLen < math.MaxInt {
		return s[minLeft:minRight]
	}
	return ""
}
