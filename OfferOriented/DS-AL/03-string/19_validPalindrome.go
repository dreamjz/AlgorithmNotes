package ordastring

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < len(s)/2 {
		if s[left] != s[right] {
			break
		}
		left++
		right--
	}

	return left == len(s)/2 || isPalindrome19(s, left+1, right) || isPalindrome19(s, left, right-1)
}

func isPalindrome19(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
