package orhash

import (
	"sort"
	"unsafe"
)

func groupAnagrams(strs []string) [][]string {
	// 哈希表
	m := make(map[string][]string, len(strs))

	result := make([][]string, 0)

	for _, s := range strs {
		// 字符串排序

		// using unsafe
		// bts := unsafe.Slice(unsafe.StringData(s), len(s)) // go 1.20
		// bts := *(*[]byte)(unsafe.Pointer(&s)) // before go 1.20
		// bs := make([]byte, len(bs))
		// copy(bs, bts)

		// not using unsafe
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		sortedStr := unsafe.String(unsafe.SliceData(bs), len(bs))
		// sortedStr := *(*string)(unsafe.Pointer(&bs)) // before go 1.20
		// 记录至哈希表
		if _, ok := m[sortedStr]; !ok {
			m[sortedStr] = make([]string, 0)
		}
		m[sortedStr] = append(m[sortedStr], s)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
