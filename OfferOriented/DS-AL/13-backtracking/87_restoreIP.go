package ofbacktarck

import "strconv"

func restoreIpAddresses(s string) []string {
	res := []string{}
	helper87(s, 0, 0, "", "", &res)
	return res
}

func helper87(s string, idx, segI int, seg, ip string, res *[]string) {
	// 恢复完成
	if idx == len(s) && segI == 3 && isValid(seg) {
		*res = append(*res, ip+seg)
		return
	}

	if idx < len(s) {
		// 添加字符
		c := s[idx]
		newSeg := seg + string(c)
		if isValid(newSeg) {
			helper87(s, idx+1, segI, newSeg, ip, res)
		}

		// 添加分隔符
		if len(seg) > 0 && segI < 3 {
			helper87(s, idx+1, segI+1, string(c), ip+seg+".", res)
		}
	}
}

func isValid(s string) bool {
	n, _ := strconv.Atoi(s)
	return n <= 255 && (s == "0" || s[0] != '0')
}
