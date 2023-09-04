package ordastring

import "unicode"

func isPalindrome(s string) bool {
	// 边界
	if len(s) == 0 {
		return true
	}

	// 双指针
	left, right := 0, len(s)-1
	for left < right {
		lc := rune(s[left])
		rc := rune(s[right])

		if !isLetterOrDigigt(lc) {
			left++
			continue
		}
		if !isLetterOrDigigt(rc) {
			right--
			continue
		}
		// to lower case
		lc = unicode.ToLower(lc)
		rc = unicode.ToLower(rc)
		if lc != rc {
			return false
		}

		left++
		right--
	}

	return true
}

func isLetterOrDigigt(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
