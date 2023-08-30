package soint

import (
	"math"
)

func divide(dividend, divisor int32) int32 {
	// 满足溢出条件
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 判断正负
	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}

	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	// 计算商
	result := divideCore(dividend, divisor)
	if negative == 1 {
		return -result
	}
	return result
}

func divideCore(dividend, divisor int32) int32 {
	var quotient int32
	quotient = 0
	for dividend <= divisor {
		u := divisor
		var v int32 = 1
		for divisor >= (math.MinInt32>>1) && dividend <= u+u {
			u += u
			v += v
		}
		quotient += v
		dividend -= u
	}
	return quotient
}
