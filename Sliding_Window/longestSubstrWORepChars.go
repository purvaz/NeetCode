package SlidingWindow

func LengthOfLongestSubstring(s string) int {

	maxLength := 0
	var strMap = make(map[byte]bool)
	i := 0
	if len(s) == 0 {
		return maxLength
	}
	for j, _ := range s {
		for strMap[s[j]] {
			delete(strMap, s[i])
			i++
		}
		strMap[s[j]] = true
		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}
	}
	return maxLength
}
