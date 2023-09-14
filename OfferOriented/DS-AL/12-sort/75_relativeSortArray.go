package ofsort

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// count num in arr1
	cnt := make([]int, 1001)
	for _, n := range arr1 {
		cnt[n]++
	}

	// rewrite arr1
	idx := 0
	// write num in arr2 first
	for _, n := range arr2 {
		for cnt[n] > 0 {
			arr1[idx] = n
			idx++
			cnt[n]--
		}
	}
	// write the rest
	for i := range cnt {
		for cnt[i] > 0 {
			arr1[idx] = i
			idx++
			cnt[i]--
		}
	}

	return arr1
}
