package DynamicProgramming

func longestPalindrome(s string) string {

	result := ""
	resultLength := 0

	for i := 0; i < len(s); i++ {

		// For odd length
		left, right := i, i

		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > resultLength {
				result = s[left : right+1]
				resultLength = right - left + 1
			}
			left--
			right++
		}

		// For even length
		left, right = i, i+1

		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > resultLength {
				result = s[left : right+1]
				resultLength = right - left + 1
			}
			left--
			right++
		}
	}
	return result
}
