package array

import (
	"strings"
)

// 方法一
// 使用语言 API
func reverseWords1(s string) string {
	strs := strings.Fields(s)
	reverseStrSlice(strs)
	return strings.Join(strs, " ")
}

func reverseStrSlice(strs []string) {
	n := len(strs)
	for i := 0; i < n/2; i++ {
		strs[i], strs[n-1-i] = strs[n-1-i], strs[i]
	}
}

// 方法二
// 去除空格并反转字符串
func reverseWords2(s string) string {
	b := trim(s)
	reverseByteSlice(b, 0, len(b)-1)
	reverseEachWord(b)
	return string(b)
}

func reverseEachWord(b []byte) {
	start, end := 0, 0

	for start < len(b) {
		for end < len(b) && b[end] != ' ' {
			end++
		}
		// reverse one word
		reverseByteSlice(b, start, end-1)
		// update start index
		start = end + 1
		end = start
	}
}

func reverseByteSlice(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func trim(s string) []byte {
	res := make([]byte, 0)

	left, right := 0, len(s)-1
	for s[left] == ' ' {
		left++
	}
	for s[right] == ' ' {
		right--
	}

	for left <= right {
		if s[left] != ' ' {
			res = append(res, s[left])
		} else if res[len(res)-1] != ' ' {
			res = append(res, s[left])
		}
		left++
	}
	return res
}

// 方法三
// 使用 双端队列
type strDeQueue []string

func (sdq strDeQueue) AddFront(str string) strDeQueue {
	rcv := []string{str}
	return append(rcv, sdq...)
}

func reverseWords3(s string) string {
	n := len(s)
	// 去除两边的空格
	left, right := 0, n-1
	for s[left] == ' ' {
		left++
	}
	for s[right] == ' ' {
		right--
	}
	// 单词入列
	var b []byte
	sdq := make(strDeQueue, 0)
	for left <= right {
		c := s[left]
		if c != ' ' {
			b = append(b, c)
		} else if len(b) != 0 {
			sdq = sdq.AddFront(string(b))
			b = b[:0]
		}
		left++
	}
	sdq = sdq.AddFront(string(b))
	return strings.Join(sdq, " ")
}
