package soint

func singleNumberOther2Times(nums []int64) int64 {
	var result int64
	for i := range nums {
		result ^= nums[i]
	}
	return result
}

func singleNumberOther3Times(nums []int64) int64 {
	buf := make([]int64, 64)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 64; j++ {
			buf[j] += (nums[i] >> (63 - j)) & 1
		}
	}
	var result int64
	for i := 0; i < 64; i++ {
		u := buf[i] % 3
		result = (result << 1) + u
	}

	return result
}

func mTimesNumberOtherNTimes(nums []int64, n int64) int64 {
	buf := make([]int64, 64)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 64; j++ {
			buf[j] += (nums[i] >> (63 - j)) & 1
		}
	}

	var result int64
	for i := 0; i < 64; i++ {
		if buf[i]%n != 0 {
			result = (result << 1) + 1
		} else {
			result = result << 1
		}
	}
	return result
}
