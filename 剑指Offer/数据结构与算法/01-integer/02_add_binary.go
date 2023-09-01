package soint

import (
	"unsafe"
)

func addBinary(a, b string) string {
	lenA, lenB := len(a), len(b)
	buf := make([]byte, max(lenA, lenB)+1)

	carry := 0
	// 从最低位开始
	for i, j := lenA-1, lenB-1; i >= 0 || j >= 0; {
		digitA, digitB := 0, 0
		if i >= 0 {
			digitA = toDigit(a[i])
		}
		if j >= 0 {
			digitB = toDigit(b[j])
		}
		// 按位计算
		digitSum := digitA + digitB + carry

		var charSum byte = '0'
		if digitSum > 1 {
			carry = 1
			if digitSum > 2 {
				charSum = '1'
			}
		}
		if digitSum == 1 {
			charSum = '1'
			carry = 0
		}
		if digitSum == 0 {
			carry = 0
		}
		// 存入结果
		buf[max(i, j)+1] = charSum

		i--
		j--
	}

	if carry > 0 {
		buf[0] = '1'
	} else {
		buf = buf[1:] // 首位为0，则抛弃
	}

	// 将字节数组转化为字符串
	result := unsafe.String(unsafe.SliceData(buf), len(buf))
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toDigit(c byte) int {
	if c-'0' == 1 {
		return 1
	}
	return 0
}
