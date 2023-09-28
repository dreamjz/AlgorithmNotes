package golang

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]bool, n)

	maxLen := 1
	start := 0
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if j-i > 2 {
				dp[j] = dp[j-1] && s[i] == s[j]
			} else {
				dp[j] = s[i] == s[j]
			}

			if dp[j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}
